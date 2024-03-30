package discord_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"slices"
	"testing"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

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

	rv.SetPublicKey(pk)
	rv.SetSignature(sg)
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

	rv.SetPublicKey(pk)
	rv.SetSignature(sg)
	rv.SetTimestamp(ts)
	rv.SetBodyContent(bc)

	isVerified := rv.Verify()

	if isVerified {
		t.Error("isVerified must be false with invalid inputs")
	}
}

func TestRequestVerifier_VerifyShouldNotPanicIfPublicKeyIsWrongSized(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Verify() must not panic")
		}
	}()

	ts := []byte("my timestamp")
	bc := []byte("my body content")
	pk := []byte("invalid public key")
	sg := []byte("invalid signature")

	rv := new(discord.RequestVerifier)

	rv.SetPublicKey(pk)
	rv.SetSignature(sg)
	rv.SetTimestamp(ts)
	rv.SetBodyContent(bc)

	isVerified := rv.Verify()

	if isVerified {
		t.Error("isVerified must be false with wrong-sized public key")
	}
}
