package assistant

import (
	"testing"
	"derbysoft.com/derbysoft-rpc-go/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
)

func TestServiceImpl_GetHistory(t *testing.T) {
	service := InitService()
	history, err := service.GetHistory()
	if err != nil {
		log.Error(err)
	}
	util.PrintDebugJson(history)
}

func TestServiceImpl_GetAllCurrencies(t *testing.T) {
	service := InitService()
	{
		currencies, err := service.GetAllCurrencies("kraken")
		if err != nil {
			log.Error(err)
		}
		util.PrintDebugJson(currencies)
	}
	{
		currencies, err := service.GetAllCurrencies("bitfinex")
		if err != nil {
			log.Error(err)
		}
		util.PrintDebugJson(currencies)
	}
}

func TestServiceImpl_Sell(t *testing.T) {
	//service := InitService()
	//{
	//	//sell 0.02ETH in kraken
	//	req := new(TradeReq)
	//	//req.Price=0.01
	//	req.Amount = 0.02
	//	req.Platform = "kraken"
	//	req.Currency = "ETHUSD"
	//	res, err := service.Sell(req)
	//	if err != nil {
	//		log.Error(err)
	//	}
	//	log.Debug(res)
	//}
	{
		//sell 0.02ETH in bitfinex
		req := new(TradeReq)
		req.Price = 0.01
		req.Amount = 0.2
		req.Platform = "bitfinex"
		req.Currency = "ethusd"
		res, err := service.Sell(req)
		if err != nil {
			log.Error(err)
		}
		log.Debug(res)
	}
}

func TestServiceImpl_Buy(t *testing.T) {
	service := InitService()
	//{
	//	//sell 0.02ETH in kraken
	//	req := new(TradeReq)
	//	//req.Price=0.01
	//	req.Amount = 0.02
	//	req.Platform = "kraken"
	//	req.Currency = "ETHUSD"
	//	res, err := service.Buy(req)
	//	if err != nil {
	//		log.Error(err)
	//	}
	//	log.Debug(res)
	//}
	{
		//sell 0.02ETH in bitfinex
		req:=new(TradeReq)
		req.Price=0.01
		req.Amount=0.2
		req.Platform="bitfinex"
		req.Currency="ETHUSD"
		res,err:=service.Buy(req)
		if err!=nil{
			log.Error(err)
		}
		log.Debug(res)
	}
}
