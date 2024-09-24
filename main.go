package main

import (
	"hmstr_token_price_bot/internal"
	"log"
	"time"
)

const delay = 5 * time.Minute
const tokenPair = "HMSTRUSDT"

func main() {
	telegramBot := internal.NewTelegramBot()
	bybit := internal.NewTokenPricer()

	ticker := time.NewTicker(delay)

	log.Println("Bot started!")

	for {
		select {
		case <-ticker.C:
			price, err := bybit.GetTokenLastPrice(tokenPair)
			if err != nil {
				log.Println(err)
				continue
			}

			err = telegramBot.SendMessageToChannel(price + "$")
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
