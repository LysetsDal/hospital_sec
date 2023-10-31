package hospital

import (
	"context"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

type IPeer interface {
	SendToPeer(context.Context, *pb.ClientMessage) (context.Context, *pb.ClientMessage)
}

type Peer struct {
	pb.ClientServer
}


func NewPeerServer() *Peer {
	return &Peer{}
}
