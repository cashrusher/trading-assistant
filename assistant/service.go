package assistant

import (
	"github.com/kraken-go-api-client"
	"github.com/cashrusher/trading-assistant/bitfinex"
)

type service interface {
	getHistory() (History, error)
	sell(tradeReq TradeReq) (TradeRes, error)
	buy(tradeReq TradeReq) (TradeRes, error)
	getAllCurrencies() (CurrenciesRes, error)
}

type serviceImpl struct {
	kraken   *krakenapi.KrakenApi
	bitfinex *bitfinex.Client
}

func (s *serviceImpl) getHistory() (History, error) {
	
}


func (s *serviceImpl) sell(tradeReq TradeReq) (TradeRes, error) {

}


func (s *serviceImpl) buy(tradeReq TradeReq) (TradeRes, error) {

}

func (s *serviceImpl) getAllCurrencies() (CurrenciesRes, error){

}
