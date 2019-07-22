package admin

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func AddPlugin(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	i, err := strconv.Atoi(args[0])
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
	}
	database.AddPluginToServer(dm.GuildID, i)
}
