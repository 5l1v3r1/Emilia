package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/discord/cmds/utils"
)

//Loop is a function to delete empty date channels
func Loop(ds *discordgo.Session) {
	for {
		for _, guild := range ds.State.Guilds {
			//channels, _ := ds.GuildChannels(guild.ID)
			voiceStates := guild.VoiceStates
			var users bool
			var channel string
			for _, voiceState := range voiceStates {
				for _, channelID := range utils.ChannelIDS {
					if channelID == voiceState.ChannelID {
						users = true
					} else {
						users = false
						channel = channelID
					}
				}
			}
			if !users {
				delID(channel)
				ds.ChannelDelete(channel)
			}
		}
		time.Sleep(time.Second * 5)
	}
}

// Delete the channel from channelid list
func delID(channel string) {
	for i := 0; i < len(utils.ChannelIDS); i++ {
		if utils.ChannelIDS[i] == channel {
			utils.ChannelIDS = append(utils.ChannelIDS[:i], utils.ChannelIDS[i+1:]...)
		}
	}
}
