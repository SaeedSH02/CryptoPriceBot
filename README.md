```markdown
# Telegram Crypto Price Bot

A simple and interactive Telegram bot built with Go that fetches real-time cryptocurrency prices. This bot allows users to check the latest prices and 24h/1h change percentages of popular cryptocurrencies such as Bitcoin, Ethereum, Solana, Dogecoin, and more.

---

## **Features**
- 🪙 Fetch real-time cryptocurrency prices.
- 📈 View 24-hour and 1-hour percentage changes.
- 🕹️ Interactive buttons for easy navigation and selection.
- 🔧 Easy to configure and extend with new cryptocurrencies.
- 🌐 Fetches data using a lightweight public API.

---

## **Table of Contents**
- [Requirements](#requirements)
- [Installation](#installation)
- [API Overview](#api-overview)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

---

## **Requirements**

To run the bot, you will need:
- **Go 1.16 or higher**
- **A Telegram bot token** (Get one by chatting with [BotFather](https://core.telegram.org/bots#botfather))
- Internet connection

---

## **Installation**

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/Telegram_CryptoPriceBot.git
cd Telegram_CryptoPriceBot
```

### 2. Set Up Your Environment Variables
Create a `.env` file in the root directory and add your Telegram bot token:

```env
BOT_TOKEN=your-telegram-bot-token-here
```

> **Note:** Make sure to keep your `.env` file secure and **never** share it publicly.

### 3. Install Dependencies
Run the following command to download required dependencies:
```bash
go mod tidy
```

### 4. Start the Bot
Launch the bot using:
```bash
go run main.go
```

---

## **API Overview**

This bot uses the [CoinLore API](https://www.coinlore.com/cryptocurrency-data-api) to fetch cryptocurrency data. The API provides the following features:

1. **Price data** in USD for various cryptocurrencies.
2. **Percentage changes** in price over the last 1 hour and 24 hours.
3. Metadata such as coin name and symbol.

### Example API Request
To fetch data for Bitcoin:
```
GET https://api.coinlore.net/api/ticker/?id=90
```

### Example API Response
```json
[
  {
    "id": "90",
    "symbol": "BTC",
    "name": "Bitcoin",
    "price_usd": "30125.36",
    "percent_change_24h": "1.89",
    "percent_change_1h": "0.18"
  }
]
```

---

## **Usage**

### 1. Interact with the Bot
- Send `/start` to the bot.
- Select a cryptocurrency using the provided buttons.
- Receive a response with the coin's details, including price, 24h change, and 1h change.

### 2. Supported Cryptocurrencies
By default, the bot supports:
- Bitcoin
- Ethereum
- Solana
- Dogecoin
- Tron
- SUI
- Shiba Inu
- TON Coin

You can easily extend this list by modifying the `main.go` file.

---

## **Contributing**

Contributions are always welcome! Here's how to get started:
1. **Fork the repository** on GitHub.
2. Create a new branch for your feature (`git checkout -b feature-name`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push your branch (`git push origin feature-name`).
5. Submit a pull request with a detailed explanation of your changes.

---

## **License**

This project is licensed under the **MIT License**. Feel free to use, modify, and distribute this project. For more details, check the [LICENSE](LICENSE) file.

---

## **Contact**
If you encounter any issues or have questions, feel free to reach out or open an issue on the [GitHub repository](https://github.com/SaeedSH02/CryptoPriceBot).
```

