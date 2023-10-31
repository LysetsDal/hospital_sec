package hospital

import (
	"context"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

type Front interface {
	SendToHospital(context.Context, *pb.HospitalMessage) (*pb.HospitalMessage) 
    SendListToHospital(context.Context, *pb.HospitalList) (*pb.HospitalList) 
}

type Peer struct {
	pb.ClientServer
}

type Hospital struct {
	pb.HospitalServer
}

func NewPeerServer() *Peer {
	return &Peer{}
}
