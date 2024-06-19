package types

import (
	"encoding/hex"
	"fmt"
	"github.com/sfyr111/blocker/crypto"
	"github.com/sfyr111/blocker/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignBlock(t *testing.T) {
	var (
		block   = util.Randomblock()
		privKey = crypto.GeneratePrivateKey()
		pubKey  = privKey.Public()
	)

	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))
}

func TestHashBlock(t *testing.T) {
	block := util.Randomblock()
	hash := HashBlock(block)
	fmt.Println(hex.EncodeToString(hash)) // block hash
	assert.Equal(t, 32, len(hash))
}
