package hospital

import (
	"fmt"
	"log"
	"net"
	"os"

	hospitalServer "github.com/LysetsDal/hospital_sec/internal/hospital"
	pb "github.com/LysetsDal/hospital_sec/proto"
	"google.golang.org/grpc"
)

var (
	host = "localhost"
	port = "5000"
)


func Main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)
	grpcServer := grpc.NewServer()
	hospitalServiceServer := hospitalServer.NewServer()

	pb.RegisterHospitalServer(grpcServer, hospitalServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}

	
}
