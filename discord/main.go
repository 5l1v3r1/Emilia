package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mxrk/Emilia/database"
	"github.com/mxrk/Emilia/discord/cmds"
	"github.com/mxrk/Emilia/discord/cmds/fun"
	"github.com/mxrk/Emilia/discord/cmds/games"
	"github.com/mxrk/Emilia/discord/cmds/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var Router = cmds.New()
	Router.RegisterCommand("reverse", fun.Reverse)
	Router.RegisterCommand("choose", utils.Choose)
	Router.RegisterCommand("ping", cmds.Ping)
	Router.RegisterCommand("rps", games.Rps)
	Router.RegisterCommand("getxp", cmds.GetXP)
	Router.RegisterCommand("coinflip", games.Coinflip)

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
