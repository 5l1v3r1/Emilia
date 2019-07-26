package admin

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func AddLogChannel(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	if emptyLogChannel(dm.GuildID) {
		database.AddLogChannel(dm.GuildID, dm.ChannelID)
		ds.ChannelMessageSend(dm.ChannelID, "Set this channel as your logging channel!")
	} else {
		//Todo maybe add a confirmation
		ds.ChannelMessageSend(dm.ChannelID, "There is already a logging channel on this guild. I will replace it with this channel.")
		database.ReplaceLogChannel(dm.GuildID, dm.ChannelID)
	}
}

func RemoveLogChannel(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	if !emptyLogChannel(dm.GuildID) {
		database.RemoveLogChannel(dm.GuildID)
	}
}

func SendToLogChannel(ds *discordgo.Session, guildID, message string) {
	ch := database.GetLogChannel(guildID)
	if ch != "" {
		ds.ChannelMessageSend(ch, message)
	}
}

func emptyLogChannel(guildID string) bool {
	ch := database.GetLogChannel(guildID)
	if ch == "" {
		return true
	}
	return false
}
