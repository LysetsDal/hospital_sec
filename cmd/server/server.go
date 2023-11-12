package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

const (
	cert = "certs/hospital.crt"
	key  = "certs/hospital.key"
)

type Hospital struct {
	pb.UnimplementedHospitalServer
	ListenAddr string

	ConnectionsMU sync.RWMutex
	Connections   map[uuid.UUID]pb.PeerClient

	SecretsArrayMU sync.Mutex
	SecretsArray   []int32
}

func NewServer(host string, port string) *Hospital {
	return &Hospital{
		ListenAddr:   fmt.Sprintf("%s:%s", host, port),
		Connections:  make(map[uuid.UUID]pb.PeerClient),
		SecretsArray: make([]int32, 0),
	}
}

func (h *Hospital) SendToHospital(ctx context.Context, in *pb.HospitalRequest) (*pb.HospitalResponse, error) {
	msg := fmt.Sprintf("Message received from: %s", in.Payload)
	res := &pb.HospitalResponse{
		From: h.ListenAddr,
		Payload: msg,
	}
	
	return res, nil
}

func (h *Hospital) MustStart() error {
	keyPair, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds((keyPair)),
	)

	pb.RegisterHospitalServer(grpcServer, h)

	lis, err := net.Listen("tcp", h.ListenAddr)
	if err != nil {
		fmt.Printf("Failed to listen on port %s\n", h.ListenAddr)
		panic(err)
	}

	fmt.Println("Starting new HospitalServer on:", h.ListenAddr)
	return grpcServer.Serve(lis)

}

func main() {
	s := NewServer("localhost", "5000")
	log.Fatal(s.MustStart())
}
