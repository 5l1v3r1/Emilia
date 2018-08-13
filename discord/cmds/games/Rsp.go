package games

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Rsp(ds *discordgo.Session, dm *discordgo.Message) {
	fmt.Println("Test")
	ds.ChannelMessageSend(dm.ChannelID, "RSP")
}
