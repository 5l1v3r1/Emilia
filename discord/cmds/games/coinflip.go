package games

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Coinflip(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	answers := []string{"heads", "tails"}
	n := rand.Intn(2)
	PlayerChoice := args[0]
	PlayerChoice = strings.ToLower(PlayerChoice)
	if PlayerChoice == "tails" || PlayerChoice == "heads" {
		ai := answers[n]
		if PlayerChoice == ai {
			ds.ChannelMessageSend(dm.ChannelID, "You won")
		} else {
			ds.ChannelMessageSend(dm.ChannelID, "You lost")
		}
	} else {
		ds.ChannelMessageSend(dm.ChannelID, "You have to select between 'heads' and 'tails'")
	}

}
