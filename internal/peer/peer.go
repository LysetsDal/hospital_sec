package hospital

import (
	"context"

	pb "github.com/LysetsDal/hospital_sec/proto"
)

type IPeer interface {
	SendToPeer(context.Context, *pb.ClientMessage) (*pb.ClientMessage, error)
	Ping(context.Context, *pb.PeerPing) (*pb.PingEcho)
}

type Peer struct {
	pb.PeerServer
}

func NewPeerServer() *Peer {
	return &Peer{}
}
