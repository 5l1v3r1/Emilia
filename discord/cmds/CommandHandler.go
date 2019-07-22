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
	//t1 := time.Now()
	msg := mc.Message
	content := msg.Content
	prefix := c.Prefix

	if mc.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(content, prefix) {
		channel, _ := s.Channel(msg.ChannelID)
		guild, _ := s.Guild(channel.GuildID)
		author := msg.Author

		// Trim prefix
		content = strings.TrimPrefix(content, prefix)

		_ = guild
		_ = author

		//	content = fields[0][len(prefix):]
		split := strings.Split(content, " ")
		cmd := split[0]
		args := split[1:]

		plugins := database.GetPluginsForGuild(mc.GuildID)
		for _, rv := range c.Routes {
			if rv.Cmd == cmd {
				id := int64(rv.ID)
				for _, i := range plugins {
					if id == i {

						rv.Run(s, mc.Message, args)
						// t2 := time.Now()
						// diff := t2.Sub(t1)
						// fmt.Println(diff)
						return
					}
				}
				s.ChannelMessageSend(mc.ChannelID, "Not available on your guild. Please activate it.")
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
