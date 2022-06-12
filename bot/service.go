package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
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
		//File:    &discordgo.File{Name: "", ContentType: "", Reader: nil},
		//Files:   []*discordgo.File{},
		//Components: []discordgo.MessageComponent{
		//	discordgo.ActionsRow{
		//		Components: []discordgo.MessageComponent{
		//			discordgo.Button{
		//				Label: "Take",
		//				Style: discordgo.PrimaryButton,
		//			},
		//		},
		//	},
		//},
	}
	msg, err := session.ChannelMessageSendComplex(s.channelId, &data)
	if err != nil {
		log.Printf("There is some error %v", err)
	}

	log.Printf("messageId: %v", msg.ID)
}
