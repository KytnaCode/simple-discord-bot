package commands

import "github.com/kytnacode/simple-discord-bot/pkg/discord"

const pingType = 4

func PingHandler() discord.InteractionReply {
	return discord.InteractionReply{
		Type: pingType,
		Data: &discord.InteractionReplyData{
			TTS:     false,
			Content: "Pong!",
		},
	}
}
