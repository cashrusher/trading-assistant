package assistant

import (
	"github.com/kraken-go-api-client"
	"github.com/cashrusher/trading-assistant/bitfinex"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"derbysoft.com/derbysoft-rpc-go/log"
)

type Service interface {
	GetHistory() (*History, error)
	Sell(tradeReq *TradeReq) (*TradeRes, error)
	Buy(tradeReq *TradeReq) (*TradeRes, error)
	GetAllCurrencies() (*CurrenciesRes, error)
}

func InitService() Service {
	kraken := krakenapi.New("RNL8qrMdKy+wRwCCR7cm5xHN09Bsew3snZIN3aW3rlnLPvtHTkCKvS+u", "WP90951w5I9uFCLabh8x0SqaKaqeTCe+orIez89Io/68R8i9Xh5lnQeSOsXtlTpf4KJ+ryf8kRMFHyRzuBpfSg==")
	bitfinex := bitfinex.NewClient().Auth("eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt", "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb")
	return &ServiceImpl{kraken: kraken, bitfinex: bitfinex}
}

type ServiceImpl struct {
	kraken   *krakenapi.KrakenApi
	bitfinex *bitfinex.Client
}

func (s *ServiceImpl) GetHistory() (*History, error) {
	//krakenTradeHistory, err := s.kraken.Trades(krakenapi.XETHZUSD, 0)
	krakenOpenOrders, err := s.kraken.OpenOrders(nil)
	if err != nil {
		return nil, err
	}
	krakenClosedOrders, err := s.kraken.ClosedOrders(nil)
	if err != nil {
		return nil, err
	}
	//https://api.bitfinex.com/v1/mytrades?symbol=ETHUSD&timestamp=0&limit_trades=32
	util.PrintDebugJson(krakenOpenOrders)
	util.PrintDebugJson(krakenClosedOrders)
	bitfinexAllOrders, err := s.bitfinex.Orders.All()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	history, err := Translate2History(krakenOpenOrders.Open, krakenClosedOrders.Closed, bitfinexAllOrders)
	util.PrintDebugJson(history)
	return nil, nil
}

func (s *ServiceImpl) Sell(tradeReq *TradeReq) (*TradeRes, error) {
	return nil, nil
}

func (s *ServiceImpl) Buy(tradeReq *TradeReq) (*TradeRes, error) {
	return nil, nil
}

func (s *ServiceImpl) GetAllCurrencies() (*CurrenciesRes, error) {
	return nil, nil
}
