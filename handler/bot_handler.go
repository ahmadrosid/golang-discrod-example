package handler

import (
	"github.com/ahmadrosid/golang-discord-example/bot"
	"github.com/bwmarrin/discordgo"
	"log"
)

type botHandler struct {
	botService bot.Service
}

func NewBotHandler(service bot.Service) *botHandler {
	return &botHandler{
		botService: service,
	}
}

func (h botHandler) OnReady(session *discordgo.Session, message *discordgo.Ready) {
	log.Println("bot is ready to listen the event")
}

func (h botHandler) OnInteraction(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("interaction created")
	h.botService.ResponseInteraction(session, interaction)
}

func (h botHandler) SendQuestionToChannel(session *discordgo.Session) {
	h.botService.SendQuestion(session)
}
