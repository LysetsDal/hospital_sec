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
	port     = flag.String("port", "5000", "port for peer")
	certFile = flag.String("cert", "certs/alice.crt", "load cert file")
	keyFile  = flag.String("key", "certs/alice.key", "load key file")
)

type Peer struct {
	pb.UnimplementedPeer2PeerServer
	ListenAddr string
	Name       string

	Peers   map[string]pb.Peer2PeerClient
	PeerDNS map[string]string

	SecretMU     sync.Mutex
	SecretShares map[string]int32
}

func NewPeer(host, port string) *Peer {
	return &Peer{
		ListenAddr:   fmt.Sprintf("%s:%s", host, port),
		Peers:        make(map[string]pb.Peer2PeerClient),
		PeerDNS:      make(map[string]string),
		SecretShares: make(map[string]int32),
	}
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

func (p *Peer) Start() error {
	p.Name = *name
	p.initPeerDNS()
	go p.StartListening(*certFile, *keyFile)
	p.readLoop()

	select {}
}

func main() {
	flag.Parse()
	p := NewPeer(host, *port)

	p.Start()
}

func (p *Peer) readLoop() {
	time.Sleep(200 * time.Millisecond)

	for {
		input := promptInput("Enter Message: \n")

		switch strings.TrimSpace(input) {
		case "sendPeer":
			name := promptInput("Enter name: \n")
			if !p.PeerIsConnected(name) {
				log.Printf("Peer {%s} not found\n", name)
				continue
			}
			text := promptInput("Enter message: \n")
			p.HandleSendMessageToPeer(name, text)
			continue

		case "broadcast":
			text := promptInput("Enter broadcast message: \n")
			p.HandleBroadcastToPeers(text)
			continue

		case "secret":
			secret, _ := promptSecretInput("Enter secret: \n")
			p.HandleInitiateSecretShare(secret)
			// Sum and send share to peers

			// reconstruct all outputs to original secret
			continue

		case "getSecrets":
			p.printSecrets()
		default:
			fmt.Println("Invalid option. Try again")
			continue
		}
	}
}

// ========== MESSAGE TO PEERS ===========
func (p *Peer) HandleBroadcastToPeers(text string) {
	for name := range p.PeerDNS {
		if name == p.Name {
			continue
		}
		p.HandleSendMessageToPeer(name, text)
	}
}

func (p *Peer) HandleSendMessageToPeer(name, text string) {
	targetAddr := p.PeerDNS[name]
	target, exists := p.Peers[targetAddr]
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
		log.Printf("Error sending message to peer %s: %v\n", targetAddr, err)
		return
	}

	fmt.Printf("Peer {%s}: Message received\n", res.FromPeer)
}

func (p *Peer) SendMessageToPeer(ctx context.Context, in *pb.PeerRequest) (*pb.PeerReply, error) {
	log.Printf("Message from {%s} - %s\n", in.FromPeer, in.Payload)
	logMessage := fmt.Sprintf("Peer {%s}: message received - %s", p.PeerDNS[p.Name], in.Payload)
	return &pb.PeerReply{
		FromPeer: p.ListenAddr,
		Payload:  logMessage,
	}, nil
}

// ======= SECRET SHAREING =========
func promptSecretInput(prompt string) (int32, error) {
	input := promptInput(prompt)
	secret, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return int32(secret), nil
}

func (p *Peer) HandleInitiateSecretShare(secret int32) {
	numPeers := int32(len(p.Peers))
	// fmt.Printf("Len p.Peers %v\n", numPeers) // DEBUGGING

	if numPeers == 0 {
		fmt.Println("No peers available to share the secret with.")
		return
	}

	// Split secret into shares:
	shares, err := util.SplitSecret(secret, numPeers)
	if err != nil {
		log.Printf("Error splitting secret into shares %v\n", err)
		return
	}
	// DEBUGGING
	fmt.Printf("=========== Spilt secrets into array: ===========\n")
	for i := range shares {
		fmt.Printf("Shares: %v\n", shares[i])
	}

	// Map name of peer to a share in the map:
	p.populateSecretsMap(shares)
	// DEBUGGING
	p.printSecrets()

	// Send a share of split secret to peers:
	p.SecretMU.Lock()
	for p_name, secret := range p.SecretShares {
		if p_name == p.Name {
			continue
		}

		targetAddr := p.PeerDNS[p_name]
		target, exists := p.Peers[targetAddr]
		if !exists {
			log.Println("Failed to find target addr in peers")
			return
		}

		msg := &pb.SecretShare{
			FromPeer: p.Name,
			Share:    secret,
		}

		res, err := target.InitiateSecretShare(context.Background(), msg)
		if err != nil {
			log.Printf("Error exchanging share with peer %s: %v\n", targetAddr, err)
			return
		}
		// INSERT THIS NEW VALUE INTO CORRESPONDING SECRETS ARRAY:
		p.SecretShares[res.FromPeer] = res.Share
		log.Printf("Peer {%s}: Share {%d} received\n", res.FromPeer, res.Share)
	}
	p.SecretMU.Unlock()
}

// Split and populate own secrets.
// send secret corresponding to in.FromPeer back.
func (p *Peer) InitiateSecretShare(ctx context.Context, in *pb.SecretShare) (*pb.SecretReply, error) {
	log.Printf("Message from {%s} - Share: %d\n", in.FromPeer, in.Share)
	shares, err := util.SplitSecret(20, int32(len(p.Peers)))
	if err != nil {
		log.Printf("Error splitting secret %v", err)
		return nil, err
	}
	fmt.Printf("=========== Spilt secrets into array: ===========\n")
	for i := range shares {
		fmt.Printf("Shares: %v\n", shares[i])
	}
	p.populateSecretsMap(shares)
	p.printSecrets()

	// Store old share from 
	oldShare := p.SecretShares[in.FromPeer]

	targetAddr := p.PeerDNS[in.FromPeer]
	_, exists := p.Peers[targetAddr]
	if !exists {
		log.Println("Failed to find target addr in peers")
	}

	p.SecretShares[in.FromPeer] = in.Share

	return &pb.SecretReply{
		FromPeer: p.Name, Payload: "", Share: oldShare}, nil
}

func (p *Peer) SendSecretToPeer(ctx context.Context, in *pb.SecretShare) (*pb.SecretReply, error) {
	log.Printf("Message from {%s} - Share: %d\n", in.FromPeer, in.Share)
	logMessage := fmt.Sprintf("Peer {%s}: Share received - %d\n", p.PeerDNS[p.Name], in.Share)

	return &pb.SecretReply{FromPeer: p.ListenAddr, Payload: logMessage}, nil
}

func (p *Peer) populateSecretsMap(shares []int32) {
	p.SecretMU.Lock()
	defer p.SecretMU.Unlock()
	// Populate the secret shares map.
	i := 0
	for p_name := range p.PeerDNS {
		// Store the share in the map using the peer's name as the key
		p.SecretShares[p_name] = shares[i]
		i++
	}
}

func (p *Peer) printSecrets() {
	fmt.Printf("=========== POPULATED MAP ===========\n")
	for name, share := range p.SecretShares {
		fmt.Printf("Name: %s - Share: %d\n", name, share)
	}
}

func (p *Peer) PeerIsConnected(name string) bool {
	_, exists := p.PeerDNS[name]
	return exists
}

func (p *Peer) AddPeerConn(ip_addr string) error {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(ip_addr,
		grpc.WithTransportCredentials(
			credentials.NewTLS(util.LoadTLSConfig(*certFile, *keyFile))))
	if err != nil {
		log.Fatalf("Could not connect to peer: %s\n", err)
		return err
	}

	log.Printf("Success on dial to Peer: %v\n", ip_addr)
	newPeer := pb.NewPeer2PeerClient(conn)
	p.Peers[ip_addr] = newPeer
	return nil
}

func (p *Peer) initPeerDNS() {

	peers := map[string]string{
		"alice":   "localhost:8080",
		"bob":     "localhost:8081",
		"charlie": "localhost:8082",
	}

	for peer, port := range peers {
		p.AddPeerConn(port)
		p.PeerDNS[peer] = port
	}
}

func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
