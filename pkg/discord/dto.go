package discord

type Dto struct {
	Type int `json:"type"`
}

type InteractionReply struct {
	Type int                   `json:"type"`
	Data *InteractionReplyData `json:"data"`
}

type InteractionReplyData struct {
	TTS     bool   `json:"tts"`
	Content string `json:"content"`
}

type InteractionRequest struct {
	Type int                     `json:"type"`
	Data *InteractionRequestData `json:"data"`
}

type InteractionRequestData struct {
	Name string `json:"name"`
}
