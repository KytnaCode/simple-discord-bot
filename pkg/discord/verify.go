package discord

import (
	"crypto/ed25519"
	"fmt"
	"slices"
)

type Verifier interface {
	Verify() bool
}

type RequestVerifier struct {
	publicKey, signature, timestamp, bodyContent []byte
}

func (rv *RequestVerifier) SetPublicKey(pk []byte) error {
	if len(pk) != ed25519.PublicKeySize {
		return fmt.Errorf("public key bad size: expected %v got %v", ed25519.PublicKeySize, len(pk))
	}

	rv.publicKey = pk

	return nil
}

func (rv *RequestVerifier) SetSignature(sg []byte) error {
	if len(sg) != ed25519.SignatureSize {
		return fmt.Errorf("signature bad size: expected %v got %v", ed25519.SignatureSize, len(sg))
	}

	rv.signature = sg

	return nil
}

func (rv *RequestVerifier) SetTimestamp(ts []byte) {
	rv.timestamp = ts
}

func (rv *RequestVerifier) SetBodyContent(bc []byte) {
	rv.bodyContent = bc
}

func (rv *RequestVerifier) Verify() bool {
	isVerified := ed25519.Verify(
		rv.publicKey,
		slices.Concat(rv.timestamp, rv.bodyContent),
		rv.signature,
	)

	return isVerified
}
