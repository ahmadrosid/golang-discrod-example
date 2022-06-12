package main

import (
	"github.com/ahmadrosid/golang-discord-example/bot"
	"github.com/ahmadrosid/golang-discord-example/config"
	"github.com/ahmadrosid/golang-discord-example/handler"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"time"
)

func initBot(cfg config.Config, f func(discordBot *discordgo.Session)) {
	discordBot, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Failed to start discordBot %v", err)
	}

	discordBot.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentGuildMessages | discordgo.IntentGuildMembers
	f(discordBot)

	err = discordBot.Open()
	if err != nil {
		log.Fatalf("Failed to open discordBot session: %v", err)
	}
	defer discordBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("shutdown discordBot")
}

func main() {
	cfg := config.Get()
	initBot(cfg, func(discordBot *discordgo.Session) {
		service := bot.NewService(cfg.QnaChannel)
		botHandler := handler.NewBotHandler(service)
		discordBot.AddHandler(botHandler.OnReady)
		discordBot.AddHandler(botHandler.OnInteraction)

		go func() {
			time.Sleep(time.Second * 1)
			botHandler.SendQuestionToChannel(discordBot)
		}()
	})
}
