package config

import (
	"log"
	"os"
)

// Config holds the bot's configuration
type Config struct {
	BotToken string
	ApiID    string
	ApiHash  string
	OwnerID  string
}

// LoadConfig loads necessary configurations from environment variables
func LoadConfig() *Config {
	cfg := &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		ApiID:    os.Getenv("API_ID"),
		ApiHash:  os.Getenv("API_HASH"),
		OwnerID:  os.Getenv("OWNER_ID"),
	}

	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is required. Obtain it from https://core.telegram.org/bots#botfather.")
	}

	if cfg.OwnerID == "" {
		log.Fatal("OWNER_ID is required. Use your Telegram user ID, retrievable via @userinfobot.")
	}

	return cfg
}