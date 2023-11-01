package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/LysetsDal/hospital_sec/config"
	"google.golang.org/grpc"

	hospitalServer "github.com/LysetsDal/hospital_sec/internal/hospital"
	pb "github.com/LysetsDal/hospital_sec/proto"
)

var (
	host = config.ServerHost
	port = config.ServerPort
	crt  = config.TLScert
	key  = config.TLSkey
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	keyPair, err := LoadTLSConfig(crt, key)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("tcp listener started on port: ", port)
	grpcServer := grpc.NewServer(grpc.Creds(keyPair))
	hospitalServiceServer := hospitalServer.NewServer(host, port)

	pb.RegisterHospitalServer(grpcServer, hospitalServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}

}
