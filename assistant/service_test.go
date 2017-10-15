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
