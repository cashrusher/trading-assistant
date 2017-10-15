package assistant

import (
	"github.com/cashrusher/trading-assistant/bitfinex"
	"github.com/kraken-go-api-client"
	"derbysoft.com/derbysoft-rpc-go/log"
	"time"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"strconv"
	"strings"
)

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
func Translate2HistoryResponse(open *krakenapi.OpenOrdersResponse, close *krakenapi.ClosedOrdersResponse, bitfinexAllOrders []bitfinex.Order) ([]History, error) {
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

func getBitfinexHistory(order bitfinex.Order) *History {
	history := new(History)
	history.Platform = "Bitfinex"
	history.OrderID = strconv.FormatInt(order.ID, 10)
	timeint, err := strconv.ParseInt(strings.TrimRight(order.Timestamp, ".0"), 10, 64)
	if err != nil {
		log.Error(err)
		history.Time = time.Now().Format("2006-01-02T15:04:05")
	} else {
		timeObject := time.Unix(timeint, 0)
		history.Time = timeObject.Format("2006-01-02 15:04:05")
	}
	volume, err := util.StringToFloat64(order.ExecutedAmount)
	if err != nil {
		log.Error(err)
	}
	history.Volume = volume
	amount, err := util.StringToFloat64(order.AvgExecutionPrice)
	if err != nil {
		log.Error(err)
	}
	history.Amount = amount
	price, err := util.StringToFloat64(order.Price)
	if err != nil {
		log.Error(err)
	}
	history.Price = price
	history.Fee = 0
	history.OrderType = order.Type
	if order.IsLive {
		history.Status = "Open"
	} else if order.IsCanceled {
		history.Status = "Close"
	} else if order.IsHidden {
		history.Status = "Hidden"
	}
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
