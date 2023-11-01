package hospital

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"github.com/LysetsDal/hospital_sec/config"

	hospitalServer "github.com/LysetsDal/hospital_sec/internal/hospital"
	pb "github.com/LysetsDal/hospital_sec/proto"
)

var (
	host = config.ServerHost
	port = config.ServerPort
	crt = config.TLScert
	key = config.TLSkey
)


func Main() {
	addr := fmt.Sprintf("%s:%s", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	keyPair, err := loadTLSConfig(crt, key)
    if err != nil {
        log.Println("error loading TLS config: ", err)
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
