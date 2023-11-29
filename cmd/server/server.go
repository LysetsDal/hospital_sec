package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	util "github.com/LysetsDal/hospital_sec/cmd/utils"
	pb "github.com/LysetsDal/hospital_sec/proto"
	"google.golang.org/grpc"
)

const (
	cert = "certs/hospital.crt"
	key  = "certs/hospital.key"
)

type HospitalServer struct {
	pb.UnimplementedHospitalServer
	ListenAddr string

	ML_DATA_MU sync.Mutex
	ML_DATA    int64
}

func NewServer(host string, port string) *HospitalServer {
	return &HospitalServer{
		ListenAddr: fmt.Sprintf("%s:%s", host, port),
		ML_DATA:    0,
	}
}

func main() {
	s := NewServer("localhost", "5000")
	log.Fatal(s.Start())
}

func (h *HospitalServer) Start() error {
	keyPair, _ := util.LoadServerTLSConfig(cert, key)

	grpcServer := grpc.NewServer(grpc.Creds((keyPair)))
	pb.RegisterHospitalServer(grpcServer, h)

	lis, err := net.Listen("tcp", h.ListenAddr)
	if err != nil {
		fmt.Printf("Failed to listen on port %s\n", h.ListenAddr)
		panic(err)
	}

	fmt.Printf("Starting new HospitalServer on: %s\n", h.ListenAddr)
	return grpcServer.Serve(lis)  // <-- This is a Blocking call 
}

// Writes the data to its 'Machine Learning database' 
// (THE VALUE IS OVERWRITTEN EACH TIME FOR DEMONSTRATION PURPOSES!)
func (h *HospitalServer) SendToHospital(ctx context.Context, in *pb.HospitalMessage) (*pb.HospitalResponse, error) {
	h.ML_DATA_MU.Lock()
	defer h.ML_DATA_MU.Unlock()
	h.ML_DATA = in.AnonymousAccumulatedData
	log.Printf("Hospital Machine-learning data set: %d", h.ML_DATA)
	// test
	secretUnModded := util.ReconstructSecret(h.ML_DATA)
	log.Printf("Hospital mod secret: %d", secretUnModded)
	return &pb.HospitalResponse{DataReceived: true}, nil
}
