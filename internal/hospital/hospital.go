package hospital

import (
	"context"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

type IHospital interface {
	SendToHospital(context.Context, *pb.HospitalMessage) (*pb.HospitalMessage) 
    SendListToHospital(context.Context, *pb.HospitalList) (*pb.HospitalList) 
}

type Hospital struct {
	pb.HospitalServer
}


func NewHospitalServer() *Hospital {
	return &Hospital{}
}