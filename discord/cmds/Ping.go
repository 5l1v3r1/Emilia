package cmds

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	ds.ChannelMessageSend(dm.ChannelID, "Pong")
}
