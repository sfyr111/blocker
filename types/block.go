package types

import (
	"crypto/sha256"
	pb "github.com/golang/protobuf/proto"
	"github.com/sfyr111/blocker/crypto"
	"github.com/sfyr111/blocker/proto"
)

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
