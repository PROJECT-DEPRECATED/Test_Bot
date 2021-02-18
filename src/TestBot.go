package main

import (
	"Test_Bot/src/commands"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const version = "0.0.2v"

var Token string

func init() {
	flag.StringVar(&Token, "token", "", "DISCORD_TOKEN")
	flag.Parse()
}

func main() {
	log.Println("Version:", version)
	session, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Fatal(err)
		return
	}

	if Token == "" {
		fmt.Println("Token: null")
		log.Fatal(err)

		return
	} else {
		fmt.Println("Token:", Token)
	}

	log.Println("get_token: Get Token access complete!")

	err = session.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	session.UpdateGameStatus(0, "Bot Runner")
	session.AddHandler(commands.PingPongListener)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	log.Println("authenticate: Connect complete!")
	log.Println("bot_manager: Bot is now running. Press CTRL-C to exit.")

	// Ctrl + C for exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}
