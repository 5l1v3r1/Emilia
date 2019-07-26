package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func (c *Commands) OnMessage(s *discordgo.Session, mc *discordgo.MessageCreate) {
	if mc.Author.ID == s.State.User.ID {
		return
	}
	if !database.IsGuildInDataBase(mc.GuildID) {
		database.InitGuild(mc.GuildID)
	}
	database.CheckUser(mc.Author.ID, mc.Author.Username)

}
