package node

import (
	"context"
	"fmt"
	"github.com/sfyr111/blocker/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	Version string
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		Version: "blocker-0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {

	p, _ := peer.FromContext(ctx)

	fmt.Printf("Received handshake from: %s: %+v\n", v, p.Addr)

	return &proto.Version{
		Version: n.Version,
		Height:  100,
	}, nil

}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("Received transaction from: ", peer)
	return &proto.Ack{}, nil
}
