package  main

import (
	"net/http"
	"github.com/cashrusher/trading-assistant/assistant"
)

func main(){
	http.HandleFunc("/assistant/trading-history", assistant.HistoryHandler)
	http.HandleFunc("/assistant/currencies/kraken", assistant.KrakenCurrencyHandler)
	http.HandleFunc("/assistant/currencies/bitfinex", assistant.BitfinexCurrencyHandler)
	http.HandleFunc("/assistant/buy", assistant.BuyHandler)
	//http.HandleFunc("/assistant/sell", sellHandler)
	http.ListenAndServe(":8080", nil)
}
