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
	//TODO check if plugin id is valid
	database.AddPluginToServer(dm.GuildID, i)
	s := fmt.Sprintf("Added plugin %d. Now your guild is allowed to use it.", i)
	ds.ChannelMessageSend(dm.ChannelID, s)
}

func RemovePlugin(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	i, err := strconv.Atoi(args[0])
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
	}
	if i == 1 {
		//not possible to remove admin commands
		return
	}
	database.RemovePlugin(dm.GuildID, i)
	s := fmt.Sprintf("Removed plugin %d. Now your guild isn't allowed to use it anymore.", i)
	ds.ChannelMessageSend(dm.ChannelID, s)
}
