package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"CryptoPriceBot/models"
)

func GetPrice(url string) string {
	response, err := http.Get(url)
		if err != nil {
			log.Println("Error fetching data from API:", err)
			return "Error fetching data. Please try again later."
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error", err)
		}
		var prices []models.Price
		err = json.Unmarshal(body, &prices)
		if err != nil {
			log.Println("Error", err)
		}

		price := prices[0]
		return fmt.Sprintf("⭕ Symbol: %s \n\n〽️ Name: %s \n\n💲Price: %s \n\n 〽️ Change 24h:  %s \n\n 📈 Change 1h:  %s",
		price.Symbol, price.Name, price.PriceUSD, price.Change24h, price.Change1h )
}
