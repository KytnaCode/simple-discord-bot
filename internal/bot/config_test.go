package bot_test

import (
	"testing"

	"github.com/kytnacode/simple-discord-bot/internal/bot"
)

func TestGetPublicKey_ShouldReturnAnEmptyStringIfTheRequiredEnvVarIsNotSet(t *testing.T) {
	t.Setenv("BOT_PUBLIC_KEY", "")

	result := bot.GetPublicKey()

	if result != "" {
		t.Errorf(
			"BOT_PUBLIC_KEY is not set but result is not an empty string: expected '%v' got '%v'",
			"",
			result,
		)
	}
}

func TestGetPublicKey_ShouldReturnAnNonEmptyStringIfTheRequiredEnvVarIsSet(t *testing.T) {
	t.Setenv("BOT_PUBLIC_KEY", "some public key")

	result := bot.GetPublicKey()

	if result == "" {
		t.Error("BOT_PUBLIC_KEY is set but `bot.GetPublicKey` return an empty value")
	}
}

func TestGetPublicKey_ShouldReturnEnvVarvalue(t *testing.T) {
	const expected = "MY_PUBLIC_KEY"

	t.Setenv("BOT_PUBLIC_KEY", expected)

	result := bot.GetPublicKey()

	if result != expected {
		t.Errorf("`bot.GetPublicKey()` not returned the expected value: expected '%v' got '%v'", expected, result)
	}
}
