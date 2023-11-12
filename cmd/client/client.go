package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	util "github.com/LysetsDal/hospital_sec/cmd/utils"
	pb "github.com/LysetsDal/hospital_sec/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	host     = "localhost"
	port     = flag.String("port", "6000", "port for peer")
	certFile = flag.String("cert", "certs/alice.crt", "load cert file")
	keyFile  = flag.String("key", "certs/alice.key", "load key file")
)

type Peer struct {
	pb.UnimplementedPeerServer
	ListenAddr string
	UUID       string

	Peers map[uuid.UUID]pb.PeerClient

	MessageCh chan *util.Message
	AddPeer   chan *Peer
	DelPeer   chan *Peer
}

func NewPeer(host string, port string) *Peer {
	return &Peer{
		ListenAddr: fmt.Sprintf("%s:%s", host, port),
		UUID:       uuid.New().String(),
		Peers:      make(map[uuid.UUID]pb.PeerClient),
		AddPeer:    make(chan *Peer, 10),
		DelPeer:    make(chan *Peer),
	}
}

func (p *Peer) MustStart(certFile string, keyFile string) error {
	keyPair, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(keyPair),
	)

	lis, err := net.Listen("tcp", p.ListenAddr)
	if err != nil {
		fmt.Printf("Failed to listen on port %s\n", p.ListenAddr)
		panic(err)
	}

	fmt.Println("Starting new Peer on:", p.ListenAddr)
	return grpcServer.Serve(lis)
}

func main() {
	flag.Parse()
	peer := NewPeer(host, *port)

	log.Fatal(peer.MustStart(*certFile, *keyFile))

}

func (p *Peer) Loop() {
	// reader := bufio.NewReader(os.Stdin)
	for {
		text := promptInput("Enter Message: \n")

		switch strings.TrimSpace(text) {
		case "send":
			text:= promptInput("Enter text: \n") 
			msg := &pb.ClientMessage{
				From: p.ListenAddr,
				Payload: text,
			}

		}
	}
}

func (p *Peer) SendToPeer(ctx context.Context, req *pb.ClientMessage) (*pb.ClientResponse){
	fmt.Printf("Received message from (%s): %s\n", req.From, req.GetPayload())
	return &pb.MessageRes{Res: "Message received"}, nil
}


func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}