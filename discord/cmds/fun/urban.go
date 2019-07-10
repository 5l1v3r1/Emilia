package fun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Response struct {
	List []struct {
		Definition  string        `json:"definition"`
		Permalink   string        `json:"permalink"`
		ThumbsUp    int           `json:"thumbs_up"`
		SoundUrls   []interface{} `json:"sound_urls"`
		Author      string        `json:"author"`
		Word        string        `json:"word"`
		Defid       int           `json:"defid"`
		CurrentVote string        `json:"current_vote"`
		WrittenOn   time.Time     `json:"written_on"`
		Example     string        `json:"example"`
		ThumbsDown  int           `json:"thumbs_down"`
	} `json:"list"`
}

var url = "http://api.urbandictionary.com/v0/define?term="

func Urban(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	word := args[0]
	response, err := http.Get(url + word)
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var object Response
	json.Unmarshal(responseData, &object)

	ret := createEmbed(object)
	ds.ChannelMessageSendEmbed(dm.ChannelID, ret)
}

func createEmbed(r Response) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: r.List[0].Author,
			URL:  r.List[0].Permalink,
		},
		Color:       0x0FE3DC, // Green
		Description: "Simple UD embed",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Definition",
				Value:  r.List[0].Definition,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Example",
				Value:  r.List[0].Example,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Dictionary",
	}
	return embed

}
