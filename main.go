package main

import (
	"log"
	"os"
	"slyfluffy/bible-bot/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables!")
	}

	bot.BotToken = os.Getenv("BOT_TOKEN") 
	bot.Run()
}