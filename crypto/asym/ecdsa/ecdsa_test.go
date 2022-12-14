package ecdsa

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/theneverse/neverse-kit/crypto"
	"github.com/stretchr/testify/require"
)

var msg = make([]byte, 961)

func TestSignR1(t *testing.T) {
	h := sha256.Sum256(msg)
	priv, err := New(crypto.ECDSA_P256)
	require.Nil(t, err)
	sign, err := priv.Sign(h[:])
	require.Nil(t, err)
	require.NotEqual(t, nil, sign)
	pub := priv.PublicKey()
	b, err := pub.Verify(h[:], sign)
	fmt.Println(err)
	require.Nil(t, err)
	require.True(t, b)
}

func TestGenerateKey(t *testing.T) {
	// Secp256k1 marshal not supported yet
	// keyK1, err := GenerateKey(Secp256k1)
	// require.Nil(t, err)
	//
	// _, err = keyK1.Bytes()
	// require.Nil(t, err)
	//
	// _, err = keyK1.Sign(msg)
	// require.Nil(t, err)

	keyR1, err := New(crypto.ECDSA_P256)
	require.Nil(t, err)

	_, err = keyR1.Bytes()
	require.Nil(t, err)

	_, err = keyR1.Sign(msg)
	require.Nil(t, err)
}

func TestPublicKey(t *testing.T) {
	privKeyK1, err := New(crypto.Secp256k1)
	require.Nil(t, err)

	pubKeyK1 := privKeyK1.PublicKey()
	pubBytes, err := pubKeyK1.Bytes()
	require.Nil(t, err)

	_, err = pubKeyK1.Address()
	require.Nil(t, err)

	restoredPubKey, err := UnmarshalPublicKey(pubBytes, crypto.Secp256k1)
	require.Nil(t, err)

	addr1, err := pubKeyK1.Address()
	require.Nil(t, err)
	addr2, err := restoredPubKey.Address()
	require.Nil(t, err)
	require.Equal(t, addr1, addr2)
}
