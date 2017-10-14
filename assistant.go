package  main

import "net/http"

func main(){
	http.HandleFunc("/assistant/trading-history", historyHandler)
	http.HandleFunc("/assistant/currencies", currencyHandler)
	http.HandleFunc("/assistant/buy", buyHandler)
	http.HandleFunc("/assistant/sell", sellHandler)
	http.ListenAndServe(":8080", nil)
}