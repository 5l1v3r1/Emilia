package utils

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func Choose(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	users := dm.Mentions
	random := rand.Intn(len(users)-1-0) + 0
	ret := users[random].Mention() + " was chosen!"
	ds.ChannelMessageSend(dm.ChannelID, ret)
}
