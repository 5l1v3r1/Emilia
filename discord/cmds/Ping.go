package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (m Commands) Ping(ds *discordgo.Session, dm *discordgo.Message) {
	fmt.Println("Test")
}
