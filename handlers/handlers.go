package handlers

import (
	"log"
	"CryptoPriceBot/models"
	"CryptoPriceBot/utils"
	tele "gopkg.in/telebot.v4"
)

func SetupHandlers(bot *tele.Bot) {
	bot.Handle("/start", func(ctx tele.Context) error {
		button := &tele.ReplyMarkup{}
		BTC_BTN := button.Text("Bitcoin")
		SOL_BTN := button.Text("Solana")
		ETH_BTN := button.Text("Ethereum")
		DOGE_BTN := button.Text("Doge Coin")
		SUI_BTN := button.Text("Sui")
		TRX_BTN := button.Text("TRON")
		SHIB_BTN := button.Text("Shiba Inu")
		TON_BTN := button.Text("TON Coin")
		More_BTN := button.Text("More...")
	
		button.Reply(
			button.Row(BTC_BTN,ETH_BTN),
			button.Row(SOL_BTN, DOGE_BTN,SUI_BTN),
			button.Row(TRX_BTN, SHIB_BTN,TON_BTN),
			button.Row(More_BTN),
		)
		return ctx.Send("Select the currency you want: ", button)
	})

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		text := ctx.Text()
		var url string

		switch text {
		case "Bitcoin":
			url = models.APIBaseURL + models.BTC_ID
		case "Solana":
			url = models.APIBaseURL + models.SOL_ID
		case "Ethereum":
			url = models.APIBaseURL + models.ETH_ID
		case "Doge Coin":
			url = models.APIBaseURL + models.DOGE_ID
		case "Sui":
			url = models.APIBaseURL + models.SUI_ID
		case "TRON":
			url = models.APIBaseURL + models.TRX_ID
		case "Shiba Inu":
			url = models.APIBaseURL + models.SHIB_ID
		case "TON Coin":
			url = models.APIBaseURL + models.TON_ID
		case "More...":
			return ctx.Send("No more yet...")
		default:
			log.Printf("Unsupported message: %s", text)
			return ctx.Send("Unsupported message. Please select a valid option.")
		}

		log.Printf("Fetching price for: %s UserName: %s", text, ctx.Sender().Username)
		return ctx.Send(utils.GetPrice(url))
	})
}