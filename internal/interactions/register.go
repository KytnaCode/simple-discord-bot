package interactions

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

type RegisterFunc func(context.Context, []byte) ([]byte, error)

func CreateRegisterFunc(appID, token string) RegisterFunc {
	return func(ctx context.Context, b []byte) ([]byte, error) {
		cb := new(discord.CommandBuilder)

		cb.SetAppID(appID)
		cb.SetToken(token)
		cb.SetJSON(b)

		return cb.RegisterJSON(ctx)
	}
}

func registerWorker(
	ctx context.Context,
	file string,
	c chan<- error,
	results chan<- string,
	reg RegisterFunc,
) {
	b, err := os.ReadFile(file)
	if err != nil {
		c <- err

		return
	}

	result, err := reg(ctx, b)
	if err != nil {
		c <- err

		return
	}
	results <- string(result)
}

func RegisterFolder(ctx context.Context, register RegisterFunc, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("path couldn't be readed: path %v: %w", dir, err)
	}

	var wg sync.WaitGroup

	results := make(chan string)
	c := make(chan error)

	for _, e := range entries {
		wg.Add(1)

		go func() {
			defer wg.Done()
			registerWorker(ctx, path.Join(dir, e.Name()), c, results, register)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	return nil
}
