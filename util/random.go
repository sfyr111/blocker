package util

import (
	randc "crypto/rand"
	"github.com/sfyr111/blocker/proto"
	"io"
	"math/rand"
	"time"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(randc.Reader, hash)
	return hash
}

func Randomblock() *proto.Block {
	header := &proto.Header{
		Version:      1,
		Height:       int32(rand.Intn(1000)),
		PreviousHash: RandomHash(),
		RootHash:     RandomHash(),
		Timestamp:    time.Now().UnixNano(),
	}

	return &proto.Block{
		Header: header,
	}
}
