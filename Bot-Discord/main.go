package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/Galangrs/Bot-Discord/Controllers"
	"github.com/Galangrs/Bot-Discord/Config"
)

func main() {
	// Load the Discord configuration
	tokenConfig := config.Token()
	
	// Create a new DiscordGo session
	dg, err := discordgo.New("Bot " + tokenConfig)
	if err != nil {
		fmt.Println("Error creating Discord session: %v", err)
	}

	// Add an event handler for the "messageCreate" event
	dg.AddHandler(command.MessageCreate)

	// Open a connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: %v", err)
	}

	fmt.Println("Bot is Running")

	// Wait for a CTRL+C signal to exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close the connection and perform cleanup before exiting
	err = dg.Close()
	if err != nil {
		fmt.Printf("Error closing connection: %v", err)
	}

	fmt.Println("Bot has been gracefully stopped.")
}
