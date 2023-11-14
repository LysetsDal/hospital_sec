package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	util "github.com/LysetsDal/hospital_sec/cmd/utils"
	pb "github.com/LysetsDal/hospital_sec/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	host     = "localhost"
	name     = flag.String("name", "", "Name of the peer")
	port     = flag.String("port", "8080", "port for peer")
	certFile = flag.String("cert", "certs/alice.crt", "load cert file")
	keyFile  = flag.String("key", "certs/alice.key", "load key file")
)

type Peer struct {
	pb.UnimplementedPeer2PeerServer
	ListenAddr string
	Name       string

	PeerDNS map[string]string
	Peers   map[string]pb.Peer2PeerClient

	SecretMU     sync.Mutex
	SecretShares map[string]int32
}

func NewPeer(host, port string) *Peer {
	return &Peer{
		ListenAddr:   fmt.Sprintf("%s:%s", host, port),
		PeerDNS:      make(map[string]string),
		Peers:        make(map[string]pb.Peer2PeerClient),
		SecretShares: make(map[string]int32),
	}
}

func main() {
	flag.Parse()
	p := NewPeer(host, *port)
	p.Start()
}

func (p *Peer) Start() error {
	p.Name = *name
	p.initPeerDNS()
	go p.StartListening(*certFile, *keyFile)
	p.readLoop()

	select {}
}

func (p *Peer) StartListening(certFile, keyFile string) {
	keyPair, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	lis, err := net.Listen("tcp", p.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(keyPair))
	pb.RegisterPeer2PeerServer(grpcServer, p)
	reflection.Register(grpcServer)

	// start listening
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("New Peer listening on: %s", p.ListenAddr)
}

func (p *Peer) readLoop() {
	p.welcomePrompt()
	time.Sleep(200 * time.Millisecond)
	for {
		input := promptInput("Enter command: \n")

		switch strings.TrimSpace(input) {
		case "sendPeer":
			name := promptInput("Enter name: \n")
			text := promptInput("Enter message: \n")
			p.HandleSendMessageToPeer(name, text)
			continue
		// SECRET SHARING
		case "secret":
			secret, _ := promptSecretInput("Enter secret: \n")
			// Sum and send share to peers
			p.HandleInitiateSecretShare(secret)
			// Add and send own output to peers
			p.HandleSendAddedOutputToPeer()
			// Reconstruct and send to Hospital
			p.handleSendToHospital(p.sumShares())
			continue
		// Print Secrets (For debugging)
		case "getSecrets":
			p.printSecrets()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try again")
			continue
		}
	}
}

// ========== MESSAGE TO PEER ===========
func (p *Peer) HandleSendMessageToPeer(name, text string) {
	target_ip := p.PeerDNS[name]
	target, exists := p.Peers[target_ip]
	if !exists {
		log.Println("Failed to find target addr in peers")
		return
	}

	msg := &pb.PeerRequest{
		FromPeer: p.ListenAddr,
		Payload:  text,
	}

	res, err := target.SendMessageToPeer(context.Background(), msg)
	if err != nil {
		log.Printf("Error sending message to peer %s: %v\n", target_ip, err)
		return
	}

	fmt.Printf("Peer {%s}: Message received\n", res.FromPeer)
}

func (p *Peer) SendMessageToPeer(ctx context.Context, in *pb.PeerRequest) (*pb.PeerReply, error) {
	log.Printf("Message from {%s} - %s\n", in.FromPeer, in.Payload)
	log_m := fmt.Sprintf("Peer {%s}: message received - %s", p.PeerDNS[p.Name], in.Payload)
	return &pb.PeerReply{
		FromPeer: p.ListenAddr,
		Payload:  log_m,
	}, nil
}

// ======= SECRET SHARE-ING =========
func (p *Peer) HandleInitiateSecretShare(secret int32) {
	num_peers := int32(len(p.Peers))
	if num_peers == 0 {
		fmt.Println("No peers available to share the secret with.")
		return
	}

	// SPLIT SECRET INTO SHARES
	shares, err := util.SplitSecret(secret, num_peers)
	if err != nil {
		log.Printf("Error splitting secret into shares %v\n", err)
		return
	}

	p.populateSecretsMap(shares)

	// SEND A SHARE OF THE SECRET TO ALL PEERS:
	p.SecretMU.Lock()
	defer p.SecretMU.Unlock()

	for p_name, secret := range p.SecretShares {
		if p_name == p.Name {
			continue // (DON'T SEND TO SELF)
		}

		target_ip := p.PeerDNS[p_name]
		target, exists := p.Peers[target_ip]
		if !exists {
			log.Println("Failed to find target address in Peer connections")
			return
		}

		msg := &pb.SecretMessage{
			FromPeer: p.Name,
			Share:    secret,
		}

		res, err := target.InitiateSecretShare(context.Background(), msg)
		if err != nil {
			log.Printf("Error exchanging share with peer %s: %v\n", target_ip, err)
			return
		}

		// INSERT THIS NEW VALUE INTO THE CORRESPONDING SECRETS ARRAY:
		p.SecretShares[res.FromPeer] = res.Share
		log.Printf("Peer {%s}: Share {%d} received\n", res.FromPeer, res.Share)
	}

}

// PERSON CALLING THIS INITIATES SECRET SHARE WITHIN PEER-2PEER CLUSTER
func (p *Peer) InitiateSecretShare(ctx context.Context, in *pb.SecretMessage) (*pb.SecretMessage, error) {
	log.Printf("Message from {%s} - Share: %d\n", in.FromPeer, in.Share)
	shares, err := util.SplitSecret(30, int32(len(p.Peers))) // <-- SECRET IS HARDCODED PT.
	if err != nil {
		log.Printf("Error splitting secret %v", err)
		return nil, err
	}

	p.populateSecretsMap(shares)

	// STORE OWN COMPUTED SHARE OF THE PERSON SENDING YOU A SHARE
	old_share := p.SecretShares[in.FromPeer]

	target_ip := p.PeerDNS[in.FromPeer]
	_, exists := p.Peers[target_ip]
	if !exists {
		log.Println("Failed to find target addr in peers")
	}

	p.SecretShares[in.FromPeer] = in.Share

	return &pb.SecretMessage{FromPeer: p.Name, Share: old_share}, nil
}

func (p *Peer) HandleSendAddedOutputToPeer() {
	share_sum := p.sumShares() // SUM AND SAVE SHARES

	out := &pb.SecretMessage{
		FromPeer: p.Name,
		Share:    share_sum,
	}

	// SEND TO ALL OTHER PEERS
	for p_name := range p.SecretShares {
		if p_name == p.Name {
			continue
		}

		target_ip := p.PeerDNS[p_name]
		target := p.Peers[target_ip]

		res, err := target.SendAddedOutputToPeer(context.Background(), out)
		if err != nil {
			log.Printf("Error exchanging share with peer %s: %v\n", target_ip, err)
			return
		}
		p.SecretShares[res.FromPeer] = res.Share // SAVE OTHER PEERS ACCUMULATED OUTPUTS
	}
}

func (p *Peer) SendAddedOutputToPeer(ctx context.Context, in *pb.SecretMessage) (*pb.SecretMessage, error) {
	log.Printf("Message from {%s} - Share: %d\n", in.FromPeer, in.Share)
	share_sum := p.sumShares()
	p.SecretShares[in.FromPeer] = in.Share

	return &pb.SecretMessage{FromPeer: p.Name, Share: share_sum}, nil
}

// =============== UTILITY FUNCTIONS ====================

// MAP THE NAME OF A PEER TO A SHARE IN THE SECRETS MAP
func (p *Peer) populateSecretsMap(shares []int32) {
	p.SecretMU.Lock()
	defer p.SecretMU.Unlock()

	i := 0
	for p_name := range p.PeerDNS {
		p.SecretShares[p_name] = shares[i]
		i++
	}
}

// SUM AND SAVE THE SHARES IN THE MAP
func (p *Peer) sumShares() int32 {
	p.SecretMU.Lock()
	defer p.SecretMU.Unlock()

	sum := int32(0)
	for _, share := range p.SecretShares {
		sum += share
		share = 0
	}

	p.SecretShares[p.Name] = sum
	return sum
}

// DIAL THE HOSPITAL AND SEND ACCUMULATED DATA
func (p *Peer) handleSendToHospital(data int32) {
	conn, err := grpc.Dial("localhost:5000",
		grpc.WithTransportCredentials(credentials.NewTLS(util.LoadTLSConfig(*certFile, *keyFile))))
	if err != nil {
		log.Fatalf("Could not connect to hospital: %s\n", err)
	}

	hospital := pb.NewHospitalClient(conn)

	out := &pb.HospitalMessage{
		AnonymousAccumulatedData: data,
	}

	res, _ := hospital.SendToHospital(context.Background(), out)

	log.Printf("Hospital received data: %v", res.DataReceived)
}

// SIMULATED DNS FOR THE PEERS (LOOKUP TABLE)
func (p *Peer) initPeerDNS() {
	peers := map[string]string{
		"alice":   "localhost:8080",
		"bob":     "localhost:8081",
		"charlie": "localhost:8082",
	}

	for p_name, ip_addr := range peers {
		p.addPeerConn(ip_addr)
		p.PeerDNS[p_name] = ip_addr
	}
}

func (p *Peer) addPeerConn(ip_addr string) error {
	conn, err := grpc.Dial(ip_addr,
		grpc.WithTransportCredentials(credentials.NewTLS(util.LoadTLSConfig(*certFile, *keyFile))))
	if err != nil {
		log.Fatalf("Could not connect to peer: %s\n", err)
		return err
	}

	log.Printf("Success on dial to Peer: %v\n", ip_addr)
	new_peer := pb.NewPeer2PeerClient(conn)
	p.Peers[ip_addr] = new_peer
	return nil
}

// ======== I/O HELPER FUNCTIONS =========
func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func promptSecretInput(prompt string) (int32, error) {
	input := promptInput(prompt)
	secret, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return int32(secret), nil
}

func (p *Peer) welcomePrompt() {
	fmt.Println()
	fmt.Println("====== COMMANDS ======")
	fmt.Println("sendPeer - Send a message to a peer, by name {alice, bob, charlie}")
	fmt.Println("secret - Initiate secret sharing")
	fmt.Println("exit - terminate program")
	fmt.Println()
}

// FOR DEBUGGING
func (p *Peer) printSecrets() {
	fmt.Printf("=========== POPULATED MAP ===========\n")
	for name, share := range p.SecretShares {
		fmt.Printf("Name: %s - Share: %d\n", name, share)
	}
}
