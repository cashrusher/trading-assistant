//package  main
//
//import (
//	"net/http"
//	"github.com/cashrusher/trading-assistant/assistant"
//)
//
//func main(){
//	http.HandleFunc("/assistant/trading-history", assistant.HistoryHandler)
//	http.HandleFunc("/assistant/currencies/kraken", assistant.KrakenCurrencyHandler)
//	http.HandleFunc("/assistant/currencies/bitfinex", assistant.BitfinexCurrencyHandler)
//	//http.HandleFunc("/assistant/buy", buyHandler)
//	//http.HandleFunc("/assistant/sell", sellHandler)
//	http.ListenAndServe(":8080", nil)
//}

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}