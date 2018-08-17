package games

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Rps game
func Rps(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	fmt.Println("args: ", args)
	answers := []string{"rock", "paper", "scissors"}

	n := rand.Intn(3)
	PlayerChoice := args[0]
	PlayerChoice = strings.ToLower(PlayerChoice)
	BotChoice := answers[n]
	var ret string
	switch {
	case PlayerChoice == BotChoice:
		ret = fmt.Sprintf("Draw, we both had %v", PlayerChoice)
	case PlayerChoice == "rock" && BotChoice == "paper":
		ret = fmt.Sprintf("I won with %v against your %v", BotChoice, PlayerChoice)
	case PlayerChoice == "rock" && BotChoice == "scissors":
		ret = fmt.Sprintf("I lost with %v against your %v", BotChoice, PlayerChoice)
	case PlayerChoice == "paper" && BotChoice == "scissors":
		ret = fmt.Sprintf("I won with %v against your %v", BotChoice, PlayerChoice)
	case PlayerChoice == "paper" && BotChoice == "rock":
		ret = fmt.Sprintf("I lost with %v against your %v", BotChoice, PlayerChoice)
	case PlayerChoice == "scissors" && BotChoice == "paper":
		ret = fmt.Sprintf("I won with %v against your %v", BotChoice, PlayerChoice)
	case PlayerChoice == "scissors" && BotChoice == "rock":
		ret = fmt.Sprintf("I lost with %v against your %v", BotChoice, PlayerChoice)

	}
	ds.ChannelMessageSend(dm.ChannelID, ret)
}
