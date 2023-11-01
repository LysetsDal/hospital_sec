package hospital

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/LysetsDal/hospital_sec/proto"
	"github.com/google/uuid"
)

type IHospital interface {
	SendToHospital(context.Context, *pb.HospitalRequest) *pb.HospitalResponse
	SendListToHospital(context.Context, *pb.HospitalListReq) *pb.HospitalListRes
}

type Hospital struct {
	pb.HospitalServer
	listenAddr string

	connectionsMU sync.RWMutex
	connections   map[uuid.UUID]string

	secretsArrayMU sync.Mutex
	secretsArray   []int32
}

func NewServer(host string, port string) *Hospital {
	return &Hospital{
		listenAddr:   fmt.Sprintf("%s:%s", host, port),
		connections:  make(map[uuid.UUID]string),
		secretsArray: make([]int32, 0),
	}
}


func Start() error {
	return nil
}
