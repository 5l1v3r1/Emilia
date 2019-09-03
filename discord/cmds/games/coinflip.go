package games

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func Coinflip(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	answers := []string{"head", "tails"}
	n := rand.Intn(2)
	if len(args) == 0 {
		ds.ChannelMessageSend(dm.ChannelID, "Tails or head?")
		return
	}
	PlayerChoice := args[0]
	PlayerChoice = strings.ToLower(PlayerChoice)
	if PlayerChoice == "tails" || PlayerChoice == "head" {
		ai := answers[n]
		if PlayerChoice == ai {
			ret := "You won! Gratulations. Added 50 coins to your wallet."
			ds.ChannelMessageSend(dm.ChannelID, ret)
			database.AddCoins(dm.Author.ID, 50)

		} else {
			ds.ChannelMessageSend(dm.ChannelID, "You lost")
		}
	} else {
		ds.ChannelMessageSend(dm.ChannelID, "You have to select between 'head' and 'tails'")
	}

}
