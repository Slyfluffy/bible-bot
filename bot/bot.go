package bot

import (
 "fmt"
 "log"
 "os"
 "os/signal"
 "strings"
 "slyfluffy/bible-bot/request"
 "github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)

	// Start this conversation
	discord.Open()
	defer discord.Close() // close session, after function termination

	// keep bot running until there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func requestVerse() string {
	scripture := request.RandomVerse()
	return fmt.Sprintf("%s\n%s", scripture.Reference, scripture.Text)
}

func help() string {
	return "Help:\n" + 
		"\t!verse: I will send a random verse from the bible"
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// The various response the bot can perform
	switch {
		case strings.Contains(message.Content, "!verse"):
			discord.ChannelMessageSend(message.ChannelID, requestVerse())
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend(message.ChannelID, help())
	}
}