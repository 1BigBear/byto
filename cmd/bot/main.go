package main

import (
	"log"
	"os"

	"github.com/mymmrac/telego"
)

// Entry point of the Byto-Bot Telegram bot
func main() {
	// Load the bot token from the environment variable
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN environment variable is required. Obtain it from https://core.telegram.org/bots#creating-a-new-bot using BotFather.")
	}

	// Create a new Telegram bot instance
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger())
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Placeholder for bot startup logic
	log.Println("Bot initialized successfully. Awaiting updates...")

	// Example: Start receiving updates
	// err = bot.Start() // Uncomment this line when handlers are implemented
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}