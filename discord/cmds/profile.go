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

//GetLevel answers with the first mentioned user or the author
func GetLevel(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	user := dm.Mentions
	var ret string
	if len(user) == 0 {
		ret = fmt.Sprintf("You're level %v", database.GetLevel(dm.Author.ID))
	} else {
		level := database.GetLevel(user[0].ID)
		if level == "" {
			ret = fmt.Sprintf("%s doesn't have a level yet.", user[0].Username)
		} else {
			ret = fmt.Sprintf("%s is level %v", user[0].Username, level)
		}

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)
}
