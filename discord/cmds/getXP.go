package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func GetXP(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	ret := fmt.Sprintf("XP for %s: %v", dm.Author.Username, database.ReturnXP(dm.Author.ID))
	ds.ChannelMessageSend(dm.ChannelID, ret)
}
