package discord

import (
	"crypto/ed25519"
	"slices"
)

type Verifier interface {
	Verify() bool
}

type RequestVerifier struct {
	publicKey, signature, timestamp, bodyContent []byte
}

func (rv *RequestVerifier) SetPublicKey(pk []byte) {
	rv.publicKey = pk
}

func (rv *RequestVerifier) SetSignature(sg []byte) {
	rv.signature = sg
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
