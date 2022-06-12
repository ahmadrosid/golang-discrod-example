package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

type Service struct {
	channelId string
}

func NewService(channelId string) Service {
	return Service{channelId: channelId}
}

func (s *Service) SendQuestion(session *discordgo.Session) {
	data := discordgo.MessageSend{
		Content: "Hi new question who is gonna to take?",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Take",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "-",
					},
				},
			},
		},
	}
	msg, err := session.ChannelMessageSendComplex(s.channelId, &data)
	if err != nil {
		log.Printf("There is some error %v", err)
	}

	log.Printf("messageId: %v", msg.ID)
}

func (s *Service) ResponseInteraction(session *discordgo.Session, in *discordgo.InteractionCreate) {
	log.Printf("%+v", in.User)
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"Thanks <@%s> we will create thread for you!",
				in.Interaction.User,
			),
		},
	}
	err := session.InteractionRespond(in.Interaction, &data)
	if err != nil {
		log.Printf("%v", err)
	}

	_, err = session.MessageThreadStart(s.channelId, in.Message.ID, fmt.Sprintf("new-thread-%v", time.Now().Unix()), 60)
	if err != nil {
		log.Printf("%v", err)
	}

	time.Sleep(time.Second * 1)
	err = session.InteractionResponseDelete(in.Interaction)
	if err != nil {
		log.Printf("failed to delete in: %v", err)
	}
}
