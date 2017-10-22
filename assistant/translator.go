package assistant

import (
	"github.com/kraken-go-api-client"
	"derbysoft.com/derbysoft-rpc-go/log"
	"time"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"strconv"
	"github.com/cashrusher/trading-assistant/bitfinex/v1"
	"github.com/cashrusher/trading-assistant/bitfinex/v2"
)

func Translate2TradeRes(addOrderRes *krakenapi.AddOrderResponse) (*TradeRes, error) {
	tradeRes := new(TradeRes)
	tradeRes.Info.TradeID = addOrderRes.TransactionIds[0]
	tradeRes.Status = "success"
	return tradeRes, nil
}

func Translate2Order(order *v1.Order) (*TradeRes, error) {
	tradeRes := new(TradeRes)
	tradeRes.Status = "success"
	tradeRes.Info.TradeID = string(order.ID)
	return tradeRes, nil
}

/**
Price         string
PriceFloat    float64
Volume        string
VolumeFloat   float64
Time          int64
Buy           bool
Sell          bool
Market        bool
Limit         bool
Miscellaneous string
*/
/**
price	[price]
amount	[decimal]
exchange	[string]
type	[string]	Sell or Buy
fee_currency	[string]	Currency you paid this trade’s fee in
fee_amount	[decimal]	Amount of fees you paid for this trade
tid	[integer]	unique identification number of the trade
order_id	[integer]	unique identification number of the parent order of the trade
*/
func Translate2HistoryResponse(open *krakenapi.OpenOrdersResponse, close *krakenapi.ClosedOrdersResponse, bitfinexAllOrders []v2.Order) ([]History, error) {
	histories := make([]History, 0)
	if open != nil {
		for id, o := range open.Open {
			history := getKrakenHistory(id, o)
			histories = append(histories, *history)
		}
	}
	if close != nil {
		for id, c := range close.Closed {
			history := getKrakenHistory(id, c)
			histories = append(histories, *history)
		}
	}
	if bitfinexAllOrders != nil {
		for _, order := range bitfinexAllOrders {
			history := getBitfinexHistory(order)
			histories = append(histories, *history)
		}
	}
	return histories, nil
}

func getBitfinexHistory(order v2.Order) *History {
	history := new(History)
	history.Platform = "Bitfinex"
	history.OrderID = strconv.FormatInt(order.ID, 10)
	timeObject := time.Unix(order.MTSCreated, 0)
	history.Time = timeObject.Format("2006-01-02 15:04:05")
	history.Volume = order.Amount
	history.Amount = order.AmountOrig
	history.Price = order.Price
	history.Fee = 0
	history.OrderType = order.Type
	history.Status = string(v2.OrderStatusActive)
	history.Pair = order.Symbol
	return history
}

func getKrakenHistory(id string, o krakenapi.Order) *History {
	history := new(History)
	history.Platform = "Kraken"
	history.OrderID = id
	timeint, err := util.Float64ToInt64(o.OpenTime)
	if err != nil {
		log.Error(err)
		history.Time = time.Now().Format("2006-01-02T15:04:05")
	} else {
		timeObject := time.Unix(timeint, 0)
		history.Time = timeObject.Format("2006-01-02 15:04:05")
	}
	history.Amount = o.Cost
	history.Volume = o.VolumeExecuted
	history.Price = o.Price
	history.Fee = o.Fee
	history.OrderType = o.Description.OrderType
	history.Pair = o.Description.AssetPair
	history.Status = o.Status
	return history
}
