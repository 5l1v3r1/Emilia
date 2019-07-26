package admin

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func ChangePrefix(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	newPrefix := args[0]
	database.ChangePrefix(dm.GuildID, newPrefix)
	ret := fmt.Sprintf("Changed current prefix to %s.", newPrefix)
	ds.ChannelMessageSend(dm.ChannelID, ret)
}

//somehow not really useful
func showPrefix(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	//NOT IMPLEMENTED
}
