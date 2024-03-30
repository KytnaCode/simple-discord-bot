package discord_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"slices"
	"testing"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func TestRequestVerifier_SetPublicKeyShouldAcceptValidPublicKey(t *testing.T) {
	t.Parallel()

	pk, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("error generating keys: %v", err)
	}

	rv := new(discord.RequestVerifier)
	err = rv.SetPublicKey(pk)

	if err != nil {
		t.Errorf("error on `RequestVerifier.SetPublicKey`: %v", err)
	}
}

func TestRequestVerifier_SetPublicKeyShouldReturnAnErrorWithWrongSizedkey(t *testing.T) {
	t.Parallel()

	pk := []byte("my invalid-length key")

	rv := new(discord.RequestVerifier)
	err := rv.SetPublicKey(pk)

	if err == nil {
		t.Error("SetPublicKey must return an error with a wrong-sized key")
	}
}

func TestRequestVerifier_SetSignatureShouldAcceptValidSignature(t *testing.T) {
	t.Parallel()

	msg := []byte("Hello World!")

	_, prk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("error on key generating: %v", err)
	}

	sg := ed25519.Sign(prk, msg)

	rv := new(discord.RequestVerifier)

	err = rv.SetSignature(sg)
	if err != nil {
		t.Errorf("SetSignature returned an error with a valid signature: %v", err)
	}
}

func TestRequestVerifier_SetSignatureShouldReturnAnErrorWithWrongSizedSignature(t *testing.T) {
	t.Parallel()

	sg := []byte("my invalid signature")

	rv := new(discord.RequestVerifier)

	err := rv.SetSignature(sg)
	if err == nil {
		t.Error("SetSignature returned a nil error with an invalid signature")
	}
}

func TestRequestVerifier_VerifyShouldReturnTrueWithValidSignatureKeyAndMessage(t *testing.T) {
	t.Parallel()

	ts := []byte("my timestamp")
	bc := []byte("my body content")

	pk, prk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("error on key generating: %v", err)
	}

	sg := ed25519.Sign(prk, slices.Concat(ts, bc))

	rv := new(discord.RequestVerifier)

	err = rv.SetPublicKey(pk)
	if err != nil {
		t.Errorf("SetPublicKey return an error with a valid key: key %v", hex.EncodeToString(pk))
	}

	err = rv.SetSignature(sg)
	if err != nil {
		t.Errorf(
			"SetSignature reuturn  an error with a valid signature: signature %v",
			hex.EncodeToString(sg),
		)
	}

	rv.SetTimestamp(ts)
	rv.SetBodyContent(bc)

	isVerified := rv.Verify()

	if !isVerified {
		t.Errorf("Verify return false with valid data: expected %v got %v", true, isVerified)
	}
}

func TestRequestVerifier_VerifyShouldReturnFalseWithInvalidInputs(t *testing.T) {
	t.Parallel()

	ts := []byte("my timestamp")
	bc := []byte("my body content")

	// We'll use valid keys but from different key pairs.
	pk, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("error on key generating: %v", err)
	}

	_, prk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("error on key generating: %v", err)
	}

	sg := ed25519.Sign(prk, slices.Concat(ts, bc))

	rv := new(discord.RequestVerifier)

	err = rv.SetPublicKey(pk)
	if err != nil {
		t.Errorf("invalid public key: %v", err)
	}

	err = rv.SetSignature(sg)
	if err != nil {
		t.Errorf("invalid public key: %v", err)
	}

	rv.SetTimestamp(ts)
	rv.SetBodyContent(bc)

	isVerified := rv.Verify()

	if isVerified {
		t.Error("isVerified must be false with invalid inputs")
	}
}
