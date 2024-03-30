package discord_test

import (
	"crypto/ed25519"
	"crypto/rand"
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
