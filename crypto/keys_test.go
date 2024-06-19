package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))

	pubKey := privKey.Public()
	assert.Equal(t, pubKeyLen, len(pubKey.Bytes()))
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg, sig))

	// Test with invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("foo"), sig))

	// Test with invalid pubKey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg, sig))
}

func TestNewPrivateKeyFromString(t *testing.T) {
	//seed := make([]byte, 32)
	//io.ReadFull(rand.Reader, seed)
	//fmt.Println(hex.EncodeToString(seed))

	var (
		seed       = "4a403c3608ef520f4a677e56f3147408e0183b43ad36a0168e6c79945a603c11"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "33b326adda5e6cad6918923de2259f787e1a7a1b"
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}
