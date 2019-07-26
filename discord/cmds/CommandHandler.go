package cmds

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

type Command struct {
	ID  int
	Cmd string
	Run HandlerFunc
}

type Commands struct {
	Routes []*Command
	Prefix string
}

// handler signature
type HandlerFunc func(*discordgo.Session, *discordgo.Message, []string)

// create route
func New() *Commands {
	c := &Commands{}
	c.Prefix = "!"
	return c
}

// messages -> commands
func (c *Commands) OnMessageC(s *discordgo.Session, mc *discordgo.MessageCreate) {

	msg := mc.Message
	content := msg.Content
	prefix := database.GetGuildPrefix(mc.GuildID)

	if mc.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(content, prefix) {

		//not needed at the moment
		//channel, _ := s.Channel(msg.ChannelID)
		// guild, _ := s.Guild(channel.GuildID)
		// author := msg.Author

		// Trim prefix
		content = strings.TrimPrefix(content, prefix)

		// _ = guild
		// _ = author

		//	content = fields[0][len(prefix):]
		split := strings.Split(content, " ")
		cmd := split[0]
		args := split[1:]

		if !database.IsGuildInDataBase(mc.GuildID) {
			database.InitGuild(mc.GuildID)
		}

		for _, rv := range c.Routes {
			if rv.Cmd == cmd {

				plugin := database.GetPluginForGuild(mc.GuildID, rv.ID)

				if plugin == -1 {
					s.ChannelMessageSend(mc.ChannelID, "Not available on your guild. Please activate it.")
					return
				}
				go rv.Run(s, mc.Message, args)
				return

			}
		}
	}
}

// register commands
func (c *Commands) RegisterCommand(cmd string, f HandlerFunc, id int) (*Command, error) {
	m := Command{}
	m.Cmd = cmd
	m.Run = f
	m.ID = id
	c.Routes = append(c.Routes, &m)
	return &m, nil
}
