package cmds

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

type Command struct {
	ID  int
	Cmd string
	Run HandlerFunc
	// using  discordgo.PermissionAdministrator
	Permissions int
}

type Commands struct {
	Routes []*Command
	Prefix string
}

// handler signature
type HandlerFunc func(*discordgo.Session, *discordgo.Message, []string)

// New will create a new command.
func New() *Commands {
	c := &Commands{}
	c.Prefix = "!"
	return c
}

// OnMessageC is the main command handler. It will match the message to the correct command.
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
				check, err := checkPerms(s, mc.GuildID, mc.Author.ID, rv.Permissions)
				if err != nil {
					fmt.Println(err)
				}
				if check {
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
}

// RegisterCommand can register commands.
func (c *Commands) RegisterCommand(cmd string, f HandlerFunc, id int, Permissions int) (*Command, error) {
	m := Command{}
	m.Cmd = cmd
	m.Run = f
	m.ID = id
	m.Permissions = Permissions
	c.Routes = append(c.Routes, &m)
	return &m, nil
}

//check if the user has enough permissions
func checkPerms(s *discordgo.Session, guildID string, userID string, permission int) (bool, error) {
	member, err := s.State.Member(guildID, userID)
	if err != nil {
		if member, err = s.GuildMember(guildID, userID); err != nil {
			return false, err
		}
	}
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}
	return false, nil
}
