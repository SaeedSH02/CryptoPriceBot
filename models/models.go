// models/models.go

package models

// Price defines the structure for cryptocurrency price data
type Price struct {
	Symbol    string `json:"symbol"`
	PriceUSD  string `json:"price_usd"`
	Name      string `json:"name"`
	Change24h string `json:"percent_change_24h"`
	Change1h  string `json:"percent_change_1h"`
}

// Crypto IDs for specific coins
const (
	BTC_ID  = "90"
	ETH_ID  = "80"
	SOL_ID  = "48543"
	DOGE_ID = "2"
	TRX_ID  = "2713"
	SUI_ID  = "93845"
	SHIB_ID = "45088"
	TON_ID  = "54683"
)

// Base URL for the API
const APIBaseURL = "https://api.coinlore.net/api/ticker/?id="
