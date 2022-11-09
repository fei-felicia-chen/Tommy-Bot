package bot 

import (
	"fmt"
	"github.com/fei-felicia-chen/myBot/config"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	BotID = u.ID
	goBot.AddHandler(msgHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	fmt.Println("Bot is running!")
}

func msgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "tom" {
		random := rand.Intn(3)
		if random == 0 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "meowww")
		} else if random == 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "meowmeow")
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "meoowww")
		}
	} else if (m.Content == "TOM") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "MAAAAAAAAAAARW")
	}
}
