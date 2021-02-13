package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const version = "0.0.1v"
var session, _ = discordgo.New()

func init() {
	flag.StringVar(&session.Token, "token", "", "DISCORD_TOKEN")
	flag.Parse()
}

func main() {
	var err error

	log.Println("Version:", version)

	if err != nil {
		log.Fatal(err)
		return
	}

	if session.Token == "" {
		fmt.Println("Token: null")
		log.Fatal(err)

		return
	} else {
		fmt.Println("Token:", session.Token)
	}

	log.Println("get_token: Get Token access complete!")

	session.AddHandler(sendTestMessage)

	err = session.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	session.UpdateGameStatus(0, "Bot Runner")

	log.Println("authenticate: Connect complete!")
	log.Println("bot_manager: Bot is now running. Press CTRL-C to exit.")

	// Ctrl + C for exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

func sendTestMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!test" {
		session.ChannelMessageSend(message.ChannelID, "Hello, World!")
	}
}