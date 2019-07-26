package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func Leadboard(ds *discordgo.Session, dm *discordgo.Message, content []string) {
	//missing server parameter
	sl := database.Leaderboard()
	var ret string
	for i := 0; i < len(sl); i++ {
		user, err := ds.GuildMember(dm.GuildID, sl[i].Userid)
		if err != nil {
			fmt.Println(err)
		}

		s := fmt.Sprintf("%d.: %s with %d coins! \n", i+1, user.User.Username, sl[i].Coins)
		ret += s

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)

}
