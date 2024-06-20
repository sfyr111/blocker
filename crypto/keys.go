package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	privKeyLen   = 64
	pubKeyLen    = 32
	seedLen      = 32
	addressLen   = 20
	signatureLen = 64
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLen {
		panic("seed length must be 32")
	}
	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)

	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(message []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, message),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	//return &PublicKey{key: p.key.Public().(ed25519.PublicKey)}

	b := make([]byte, pubKeyLen)
	fmt.Printf("Before copy: %x\n", b)
	copy(b, p.key[32:])
	fmt.Printf("After copy: %x\n", b)

	return &PublicKey{
		key: b,
	}
}

func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func PublicKeyFromBytes(b []byte) *PublicKey {
	if len(b) != pubKeyLen {
		panic("public key length must be 32")
	}
	return &PublicKey{key: ed25519.PublicKey(b)}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func SignatureFromBytes(b []byte) *Signature {
	if len(b) != signatureLen {
		panic("signature length must be 64")
	}
	return &Signature{
		value: b,
	}
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pubKey *PublicKey, message []byte) bool {
	return ed25519.Verify(pubKey.key, message, s.value)
}

type Address struct {
	value []byte
}

func (a Address) Bytes() []byte {
	return a.value
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
