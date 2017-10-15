package  main

import (
	"net/http"
	"github.com/cashrusher/trading-assistant/assistant/handler"
)

func main(){
	http.HandleFunc("/assistant/trading-history", handler.HistoryHandler)
	//http.HandleFunc("/assistant/currencies", currencyHandler)
	//http.HandleFunc("/assistant/buy", buyHandler)
	//http.HandleFunc("/assistant/sell", sellHandler)
	http.ListenAndServe(":8080", nil)
}