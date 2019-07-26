package admin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mxrk/Emilia/database"
)

func ReportPerson(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	persons := dm.Mentions
	person := persons[0].ID
	reportType := args[1]
	msg := strings.Join(args[2:], " ")
	id := database.AddReport(dm.GuildID, person, dm.Author.ID, msg, getReportType(reportType))
	createEmbedSingleReport(ds, dm.GuildID, person, dm.Author.Username, reportType, msg, id)
}

func Reports(ds *discordgo.Session, dm *discordgo.Message, args []string) {
	person := dm.Mentions
	user, err := ds.User(person[0].ID)
	if err != nil {
		fmt.Println("Could not find user")
	}
	reports := database.GetReports(dm.GuildID, user.ID)
	createEmbedReports(ds, database.GetLogChannel(dm.GuildID), user.Username, reports)
}

func createEmbedSingleReport(ds *discordgo.Session, guild, person, mod, reportType, msg string, id int) {
	s := strconv.Itoa(id)
	user, err := ds.User(person)
	if err != nil {
		fmt.Println("Could not find user")
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Victim",
				Value:  user.Username,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Moderator",
				Value:  mod,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Type",
				Value:  reportType,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Reason",
				Value:  msg,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Report " + s,
	}
	ds.ChannelMessageSendEmbed(database.GetLogChannel(guild), embed)
}

func createEmbedReports(ds *discordgo.Session, guild, user string, reports []database.Report) {
	embed := &discordgo.MessageEmbed{
		Author:    &discordgo.MessageEmbedAuthor{},
		Color:     0x00ff00,
		Fields:    []*discordgo.MessageEmbedField{},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Reports for " + user,
	}

	for _, report := range reports {
		s := strconv.Itoa(report.ID)
		val := fmt.Sprintf("Victim: %s \n Mod: %s \n Type: %s \n Reason: %s", report.Victim, report.Mod, getTypeFromReport(report.ReportType), report.Msg)
		d := &discordgo.MessageEmbedField{
			Name:   s,
			Value:  val,
			Inline: true,
		}
		embed.Fields = append(embed.Fields, d)
	}
	ds.ChannelMessageSendEmbed(guild, embed)
}
func getReportType(t string) int {
	switch t {
	case "warn":
		return 0
	case "ban":
		return 1
	default:
		return -1
	}
}

func getTypeFromReport(i int) string {
	switch i {
	case 0:
		return "warn"
	case 1:
		return "ban"
	default:
		return ""
	}
}
