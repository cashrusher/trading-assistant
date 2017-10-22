package assistant

import (
	"encoding/json"
	"net/http"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/log"
	"io/ioutil"
	"errors"
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

/**
{
    "platform":"kraken", //can be 'kraken' or 'bitfinex'
    "currency":"BTH", // it's the currency returned by the currencies api
    "amount":34.9982,
    "price":887.000293
}
*/
func BuyHandler(w http.ResponseWriter, r *http.Request) {
	request, err := getRequest(r)
	if err != nil {
		w.Write(createErrorResponse(err))
		return
	}
	treq := new(TradeReq)
	if err := json.Unmarshal(request, treq); err != nil {
		w.Write(createErrorResponse(err))
		return
	}
	tres, err := service.Buy(treq)
	if err != nil {
		w.Write(createErrorResponse(err))
		return
	}
	bs,err:=json.Marshal(tres)
	if err!=nil{
		w.Write(createErrorResponse(err))
	}
	w.Write(bs)
}

func getRequest(req *http.Request) ([]byte, error) {
	if err := req.ParseForm(); err != nil {
		log.Error(err)
		return nil, errors.New("Unsupport Protocol Type!")
	}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil || len(bs) == 0 {
		log.Error(err)
		return nil, errors.New("Unsupport Protocol Type!")
	}
	return bs, nil
}

