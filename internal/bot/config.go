package bot

import (
	"os"
)

func GetPublicKey() string {
	return os.Getenv("BOT_PUBLIC_KEY")
}
