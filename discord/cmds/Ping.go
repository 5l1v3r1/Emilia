package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(ds *discordgo.Session, dm *discordgo.Message) {
	fmt.Println("Test")
	ds.ChannelMessageSend(dm.ChannelID, "Pong")
}
