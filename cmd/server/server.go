package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/LysetsDal/hospital_sec/proto"

	util "github.com/LysetsDal/hospital_sec/cmd/utils"
)

const (
	cert = "certs/hospital.crt"
	key  = "certs/hospital.key"
)

type HospitalServer struct {
	pb.UnimplementedHospitalServer
	ListenAddr string

	ConnectionsMU sync.RWMutex
	Connections   []*grpc.ClientConn

	SecretsArrayMU sync.Mutex
	SecretsArray   []int32
}

func NewServer(host string, port string) *HospitalServer {
	return &HospitalServer{
		ListenAddr:   fmt.Sprintf("%s:%s", host, port),
		Connections:  make([]*grpc.ClientConn, 3),
		SecretsArray: make([]int32, 0),
	}
}

func (h *HospitalServer) MustStart() error {
	keyPair, _ := util.LoadServerTLSConfig(cert, key)

	grpcServer := grpc.NewServer(grpc.Creds((keyPair)))
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
