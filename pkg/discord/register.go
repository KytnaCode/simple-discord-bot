package discord

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

type CommandBuilder struct {
	appID string
	token string
	data  []byte
	url   string
}

func getDefaultRegisterURL(appID string) string {
	return fmt.Sprintf("https://discord.com/api/v10/applications/%v/commands", appID)
}

func makeRequest(ctx context.Context, cb *CommandBuilder) (*http.Request, error) {
	url := cb.url
	if url == "" {
		url = getDefaultRegisterURL(cb.appID)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		url,
		bytes.NewBuffer(cb.data),
	)
	if err != nil {
		return nil, fmt.Errorf("request couldn't be created: %w", err)
	}

	req.Header.Set("Authorization", "Bot "+cb.token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (cb *CommandBuilder) RegisterJSON(ctx context.Context) ([]byte, error) {
	req, err := makeRequest(ctx, cb)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("body couldn't be readed: %w", err)
	}

	return b, nil
}

func (cb *CommandBuilder) SetAppID(appID string) {
	cb.appID = appID
}

func (cb *CommandBuilder) SetToken(token string) {
	cb.token = token
}

func (cb *CommandBuilder) SetJSON(json []byte) {
	cb.data = json
}

func (cb *CommandBuilder) SetURL(url string) {
	cb.url = url
}
