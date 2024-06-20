package types

import (
	"github.com/sfyr111/blocker/crypto"
	"github.com/sfyr111/blocker/proto"
	"github.com/sfyr111/blocker/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// my balance is 100 want to send 5 coins to Alice
// 2 outputs 5 to the dude address and 95 to myself
func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	//fmt.Printf("tx: %v\n", tx)
	assert.True(t, VerifyTransaction(tx))
}
