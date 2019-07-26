package fun

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Reverse is a simple command to reverse a string
func Reverse(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	ret := strings.Join(args, " ")
	ds.ChannelMessageSend(dm.ChannelID, rev(ret))
}

func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
