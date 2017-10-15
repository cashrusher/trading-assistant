package assistant

import (
	"encoding/json"
	"net/http"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/log"
)

var service Service

func init() {
	service = InitService()
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := service.GetHistory()
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	h, err := json.Marshal(history)
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	w.Write(h)
}

func createErrorResponse(err error) []byte {
	return []byte(`{
		"status":"failed",
		"message":"` + err.Error() + `"
	}`)
}

func KrakenCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	currencies, err := service.GetAllCurrencies("kraken")
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	h, err := json.Marshal(currencies)
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	w.Write(h)
}


func BitfinexCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	currencies, err := service.GetAllCurrencies("bitfinex")
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	h, err := json.Marshal(currencies)
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	w.Write(h)
}