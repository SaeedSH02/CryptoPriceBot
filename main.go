package main


import (
	"log"
	"os"
	"time"
	"CryptoPriceBot/handlers"
	tele "gopkg.in/telebot.v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("BOT_TOKEN")

	// Bot settings
	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Setup handlers
	handlers.SetupHandlers(bot)

	// Start bot
	log.Println("Bot started...")
	bot.Start()

}