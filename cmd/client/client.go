package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	util "github.com/LysetsDal/hospital_sec/cmd/utils"
	"github.com/LysetsDal/hospital_sec/config"
	pb "github.com/LysetsDal/hospital_sec/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	host       = config.Host
	port       = config.ServerPort
	serverPort = config.ServerPort
	certFile   = config.CertFile
	keyFile    = config.KeyFile
)

type Peer struct {
	pb.PeerServer
	ListenAddr string
	Listener   net.Listener
	OwnUUID    string
	Conn       net.Conn

	PeerMU  sync.RWMutex
	Peers   map[uuid.UUID]pb.PeerClient
	AddPeer chan *Peer
	DelPeer chan *Peer

	// SecretsArray []int32
}

func NewPeer(host string, port string) *Peer {
	return &Peer{
		ListenAddr: fmt.Sprintf("%s:%s", host, port),
		OwnUUID:    uuid.New().String(),
		Peers:      make(map[uuid.UUID]pb.PeerClient),
		AddPeer:    make(chan *Peer, 10),
		DelPeer:    make(chan *Peer),


	}
}

func main() {
	flag.Parse()
	config := util.LoadTLSConfig(*certFile, *keyFile)

	peer := NewPeer(fmt.Sprint(host), port)

	log.Fatal(peer.MustStart(config))
}

func (p *Peer) MustStart(config *tls.Config) error {
	// Dial hospital server
	hospitalConn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, serverPort),
		grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		fmt.Printf("Failed to connect to server: %s\n", err)
	}
	defer hospitalConn.Close()

	hospital := pb.NewHospitalClient(hospitalConn)

}

func (p *Peer) SendToPeer(ctx context.Context, in *pb.ClientMessage) (*pb.ClientMessage, error) {
	return nil, nil
}

func (p *Peer) Ping(ctx context.Context, in *pb.PeerPing) *pb.PeerPing {
	return nil
}

func (p *Peer) getPeerList() []pb.PeerClient {
	p.PeerMU.Lock()
	defer p.PeerMU.Unlock()

	var (
		peers = make([]pb.PeerClient, len(p.Peers))
		i     = 0
	)
	for _, peer := range p.Peers {
		peers[i] = peer
		i++
	}
	return peers
}

func (p *Peer) makeGrpcClientConn(addr string, config *tls.Config) (*pb.PeerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		return nil, err
	}
	client := pb.NewPeerClient(conn)
	return &client, nil
}

type TCPTransport struct {
	listenAddr string
	listener   net.Listener
	AddPeer    chan *Peer
	DelPeer    chan *Peer
}

func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: addr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}

	t.listener = ln

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		peer := &Peer{
			Conn: conn,
		}

		t.AddPeer <- peer
	}
}
