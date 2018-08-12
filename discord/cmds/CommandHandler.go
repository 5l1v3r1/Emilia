package cmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Cmd string
	Run HandlerFunc
}

type Commands struct {
	Routes []*Command
	Prefix string
}

// handler signature
type HandlerFunc func(*discordgo.Session, *discordgo.Message)

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
		fields := strings.Fields(content)
		//	fmt.Println(fields)
		_ = guild
		_ = author

		for fk, fv := range fields {
			fmt.Println(fk, fv)
			for _, rv := range c.Routes {
				if rv.Cmd == fv {
					rv.Run(s, mc.Message)
				}
				//	fmt.Println("Range - Routes")
				//fmt.Println(*rv)
				//

			}

		}

	}

}

// register commands
func (c *Commands) RegisterCommand(cmd string, f HandlerFunc) (*Command, error) {
	m := Command{}
	m.Cmd = cmd
	m.Run = f
	c.Routes = append(c.Routes, &m)
	return &m, nil
}
