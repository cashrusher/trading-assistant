package assistant

import (
	"github.com/kraken-go-api-client"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"derbysoft.com/derbysoft-rpc-go/log"
	"strings"
	"errors"
	"github.com/cashrusher/trading-assistant/bitfinex/v2"
	"github.com/cashrusher/trading-assistant/bitfinex/v1"
)

type Service interface {
	GetHistory() ([]History, error)
	Sell(tradeReq *TradeReq) (*TradeRes, error)
	Buy(tradeReq *TradeReq) (*TradeRes, error)
	GetAllCurrencies(platform string) (CurrenciesRes, error)
}

func InitService() Service {
	kraken := krakenapi.New("RNL8qrMdKy+wRwCCR7cm5xHN09Bsew3snZIN3aW3rlnLPvtHTkCKvS+u", "WP90951w5I9uFCLabh8x0SqaKaqeTCe+orIez89Io/68R8i9Xh5lnQeSOsXtlTpf4KJ+ryf8kRMFHyRzuBpfSg==")
	bitfinexV2 := v2.NewClient().Credentials("eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt", "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb")
	bitfinexV1 := v1.NewClient().Auth("eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt", "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb")
	return &ServiceImpl{kraken: kraken, bitfinexV1: bitfinexV1, bitfinexV2: bitfinexV2}
}

type ServiceImpl struct {
	kraken     *krakenapi.KrakenApi
	bitfinexV1 *v1.Client
	bitfinexV2 *v2.Client
}

func (s *ServiceImpl) GetHistory() ([]History, error) {
	//krakenTradeHistory, err := s.kraken.Trades(krakenapi.XETHZUSD, 0)
	krakenOpenOrders, err := s.kraken.OpenOrders(nil)
	if err != nil {
		log.Error(err)
		//return nil, err
	}
	krakenClosedOrders, err := s.kraken.ClosedOrders(nil)
	if err != nil {
		log.Error(err)
		//return nil, err
	}
	//https://api.bitfinex.com/v1/mytrades?symbol=ETHUSD&timestamp=0&limit_trades=32
	util.PrintDebugJson(krakenOpenOrders)
	util.PrintDebugJson(krakenClosedOrders)
	bitfinexAllOrders, err := s.bitfinexV2.Orders.History("")
	if err != nil {
		log.Error(err)
		//return nil, err
	}
	log.Debug(bitfinexAllOrders)
	util.PrintDebugJson(bitfinexAllOrders)
	histories, err := Translate2HistoryResponse(krakenOpenOrders, krakenClosedOrders, bitfinexAllOrders)
	util.PrintDebugJson(histories)
	return histories, nil
}

func (s *ServiceImpl) Sell(tradeReq *TradeReq) (*TradeRes, error) {
	return s.trade(tradeReq, "sell")
}

func (s *ServiceImpl) Buy(tradeReq *TradeReq) (*TradeRes, error) {
	return s.trade(tradeReq, "buy")
}

func (s *ServiceImpl) trade(tradeReq *TradeReq, sellOrBuy string) (*TradeRes, error) {
	if tradeReq.Platform == "kraken" {
		volume := util.Float64ToString(tradeReq.Amount)
		price := util.Float64ToString(tradeReq.Price)
		args := make(map[string]string)
		args["price"] = price
		addOrderResponse, err := s.kraken.AddOrder(tradeReq.Currency, sellOrBuy, "market", volume, args)
		if err != nil {
			return nil, err
		}
		return Translate2TradeRes(addOrderResponse)
	} else if tradeReq.Platform == "bitfinex" {
		if sellOrBuy == "sell" {
			tradeReq.Amount = float64(0) - tradeReq.Amount
		}
		order, err := s.bitfinexV1.Orders.Create(tradeReq.Currency, tradeReq.Amount, tradeReq.Price, v1.OrderTypeMarket)

		//order,err:=s.bitfinexV1.Orders.Create(amount,price,tradeReq.Currency)
		if err != nil {
			return nil, err
		}
		return Translate2Order(order)
	}
	return nil, errors.New("Unsupported action! ")
}

func (s *ServiceImpl) GetAllCurrencies(platform string) (CurrenciesRes, error) {
	currencies := make([]string, 0)
	if platform == "kraken" {

		krakenPairs, err := s.kraken.AssetPairs()
		if err != nil {
			log.Error(err)
		}
		currencies = append(currencies, krakenPairs.BCHEUR.Altname)
		currencies = append(currencies, krakenPairs.DASHEUR.Altname)
		currencies = append(currencies, krakenPairs.DASHUSD.Altname)
		currencies = append(currencies, krakenPairs.DASHXBT.Altname)
		currencies = append(currencies, krakenPairs.EOSETH.Altname)
		currencies = append(currencies, krakenPairs.EOSEUR.Altname)
		currencies = append(currencies, krakenPairs.EOSUSD.Altname)
		currencies = append(currencies, krakenPairs.EOSXBT.Altname)
		currencies = append(currencies, krakenPairs.GNOETH.Altname)
		currencies = append(currencies, krakenPairs.GNOEUR.Altname)
		currencies = append(currencies, krakenPairs.GNOUSD.Altname)
		currencies = append(currencies, krakenPairs.GNOXBT.Altname)
		currencies = append(currencies, krakenPairs.USDTZUSD.Altname)
		currencies = append(currencies, krakenPairs.XETCXETH.Altname)
		currencies = append(currencies, krakenPairs.XETCXXBT.Altname)
		currencies = append(currencies, krakenPairs.XETCZEUR.Altname)
		currencies = append(currencies, krakenPairs.XETCXUSD.Altname)
		currencies = append(currencies, krakenPairs.XETHXXBT.Altname)
		currencies = append(currencies, krakenPairs.XETHZCAD.Altname)
		currencies = append(currencies, krakenPairs.XETHZEUR.Altname)
		currencies = append(currencies, krakenPairs.XETHZGBP.Altname)
		currencies = append(currencies, krakenPairs.XETHZJPY.Altname)
		currencies = append(currencies, krakenPairs.XETHZUSD.Altname)
		currencies = append(currencies, krakenPairs.XICNXETH.Altname)
		currencies = append(currencies, krakenPairs.XICNXXBT.Altname)
		currencies = append(currencies, krakenPairs.XLTCXXBT.Altname)
		currencies = append(currencies, krakenPairs.XLTCZEUR.Altname)
		currencies = append(currencies, krakenPairs.XLTCZUSD.Altname)
		currencies = append(currencies, krakenPairs.XMLNXETH.Altname)
		currencies = append(currencies, krakenPairs.XMLNXXBT.Altname)
		currencies = append(currencies, krakenPairs.XREPXETH.Altname)
		currencies = append(currencies, krakenPairs.XREPXXBT.Altname)
		currencies = append(currencies, krakenPairs.XREPZEUR.Altname)
		currencies = append(currencies, krakenPairs.XREPZUSD.Altname)
		currencies = append(currencies, krakenPairs.XXBTZCAD.Altname)
		currencies = append(currencies, krakenPairs.XXBTZEUR.Altname)
		currencies = append(currencies, krakenPairs.XXBTZGBP.Altname)
		currencies = append(currencies, krakenPairs.XXBTZJPY.Altname)
		currencies = append(currencies, krakenPairs.XXBTZUSD.Altname)
		currencies = append(currencies, krakenPairs.XXDGXXBT.Altname)
		currencies = append(currencies, krakenPairs.XXLMXXBT.Altname)
		currencies = append(currencies, krakenPairs.XXLMZEUR.Altname)
		currencies = append(currencies, krakenPairs.XXLMZUSD.Altname)
		currencies = append(currencies, krakenPairs.XXMRXXBT.Altname)
		currencies = append(currencies, krakenPairs.XXMRZEUR.Altname)
		currencies = append(currencies, krakenPairs.XXMRZUSD.Altname)
		currencies = append(currencies, krakenPairs.XXRPXXBT.Altname)
		currencies = append(currencies, krakenPairs.XXRPZCAD.Altname)
		currencies = append(currencies, krakenPairs.XXRPZEUR.Altname)
		currencies = append(currencies, krakenPairs.XXRPZJPY.Altname)
		currencies = append(currencies, krakenPairs.XXRPZUSD.Altname)
		currencies = append(currencies, krakenPairs.XZECXXBT.Altname)
		currencies = append(currencies, krakenPairs.XZECZEUR.Altname)
		currencies = append(currencies, krakenPairs.XZECZUSD.Altname)
		return deleteEmptyCurrency(currencies), nil
	} else if platform == "bitfinex" {
		bitfinexPairs, err := s.bitfinexV1.Pairs.All()
		if err != nil {
			log.Error(err)
		}
		currencies = append(currencies, bitfinexPairs...)
		util.PrintDebugJson(bitfinexPairs)
		return deleteEmptyCurrency(currencies), nil
	} else {
		return []string{}, nil
	}

}

func deleteEmptyCurrency(currencies []string) []string {
	result := make([]string, 0)
	for _, currency := range currencies {
		if strings.TrimSpace(currency) == "" {
			continue
		}
		result = append(result, currency)
	}
	return result
}
