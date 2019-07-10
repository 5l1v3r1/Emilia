package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var ChannelIDS []string

func Date(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	users := dm.Mentions
	st, err := ds.GuildChannelCreate(dm.GuildID, "privateRoom", discordgo.ChannelTypeGuildVoice)
	if err != nil {
		fmt.Println(err)
	}

	st.UserLimit = len(users) + 1
	for i := 0; i < len(users); i++ {
		ds.GuildMemberMove(dm.GuildID, users[i].ID, st.ID)
	}
	ChannelIDS = append(ChannelIDS, st.ID)

}
