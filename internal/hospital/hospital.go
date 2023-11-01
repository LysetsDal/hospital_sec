package hospital

import (
	"context"
	"sync"

	pb "github.com/LysetsDal/hospital_sec/proto"
	"github.com/google/uuid"
)

const (
	host = "localhost"
	port = ":8081"
)

type IHospital interface {
	SendToHospital(context.Context, *pb.HospitalRequest) *pb.HospitalResponse
	SendListToHospital(context.Context, *pb.HospitalListReq) *pb.HospitalListRes
}

type Hospital struct {
	pb.HospitalServer
	listenAddr string

	connectionsMU sync.Mutex
	connections   map[uuid.UUID]string

	secretsArrayMU sync.Mutex
	secretsArray   []int32
}

func NewServer() *Hospital {
	return &Hospital{
		listenAddr:   host + port,
		connections:  make(map[uuid.UUID]string),
		secretsArray: make([]int32, 0),
	}
}

func Start() error {
	
}