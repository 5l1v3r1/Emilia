package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mxrk/Emilia/database"
	"github.com/mxrk/Emilia/discord/cmds"
	"github.com/mxrk/Emilia/discord/cmds/admin"
	"github.com/mxrk/Emilia/discord/cmds/fun"
	"github.com/mxrk/Emilia/discord/cmds/games"
	"github.com/mxrk/Emilia/discord/cmds/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var Router = cmds.New()
	Router.RegisterCommand("reverse", fun.Reverse, 2)
	Router.RegisterCommand("urban", fun.Urban, 3)

	Router.RegisterCommand("choose", utils.Choose, 4)
	Router.RegisterCommand("date", utils.Date, 5)
	Router.RegisterCommand("leaderboard", utils.Leadboard, 6)

	Router.RegisterCommand("ping", cmds.Ping, 6)
	Router.RegisterCommand("getxp", cmds.GetXP, 7)
	Router.RegisterCommand("lvl", cmds.GetLevel, 8)
	Router.RegisterCommand("coins", cmds.GetCoins, 9)

	Router.RegisterCommand("rps", games.Rps, 10)
	Router.RegisterCommand("coinflip", games.Coinflip, 11)

	Router.RegisterCommand("addPlugin", admin.AddPlugin, 1)
	Router.RegisterCommand("removePlugin", admin.RemovePlugin, 1)
	Router.RegisterCommand("addLog", admin.AddLogChannel, 1)
	Router.RegisterCommand("removeLog", admin.RemoveLogChannel, 1)
	Router.RegisterCommand("changePrefix", admin.ChangePrefix, 1)
	Router.RegisterCommand("rep", admin.ReportPerson, 1)
	Router.RegisterCommand("reps", admin.Reports, 1)
	database.AddGame("rps", "coinflip")

	dg, err := discordgo.New("Bot " + os.Getenv("goDiscord"))
	if err != nil {
		fmt.Println(err)
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	go Loop(dg)
	//ids = make([]int, 0, 0)
	// CommandHandler
	dg.AddHandler(Router.OnMessageC)
	// MessageHandler
	dg.AddHandler(Router.OnMessage)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
