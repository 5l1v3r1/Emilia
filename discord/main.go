package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mxrk/Emilia/discord/cmds"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var Router = cmds.New()
	Router.RegisterCommand("ping", Router.Ping)

	dg, err := discordgo.New("Bot " + os.Getenv("goDiscord"))
	if err != nil {
		fmt.Println(err)
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	dg.AddHandler(Router.OnMessageC)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
