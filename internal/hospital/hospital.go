package hospital

import (
	"context"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

type IHospital interface {
	SendToHospital(context.Context, *pb.HospitalRequest) (*pb.HospitalResponse) 
    SendListToHospital(context.Context, *pb.HospitalListReq) (*pb.HospitalListRes) 
}

type Hospital struct {
	pb.HospitalServer
}


func NewServer() *Hospital {
	return &Hospital{}
}