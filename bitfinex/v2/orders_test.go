package v2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"derbysoft.com/derbysoft-rpc-go/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"strings"
)


func TestOrderService_History(t *testing.T) {
	bitfinex := NewClient().Credentials("eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt", "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb")
	orders,err:=bitfinex.Orders.History(TradingPrefix+ETHUSD)
	if err!=nil{
		log.Error(err)
	}
	util.PrintDebugJson(orders)
}

func TestOrdersAll(t *testing.T) {
	httpDo = func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
				[
					[4419360502,null,83283216761,"tIOTBTC",1508281683000,1508281731000,63938,63938,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.0000843,0,0,0,null,null,null,0,0,null],
					[4419354239,null,83265164211,"tIOTBTC",1508281665000,1508281674000,63976,63976,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.00008425,0,0,0,null,null,null,0,0,null],
					[4419339620,null,83217673277,"tIOTBTC",1508281618000,1508281653000,64014,64014,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.0000842,0,0,0,null,null,null,0,0,null]
				]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	orders, err := NewClient().Orders.All("")

	if err != nil {
		t.Error(err)
	}

	if len(orders) != 3 {
		t.Fatalf("expected three orders but got %d", len(orders))
	}
}
func currencyPairToSymbol(currencyPair CurrencyPair) string {
return strings.ToUpper(currencyPair.ToSymbol(""))
}

func  placeOrder(orderType, side, amount, price string, pair CurrencyPair) (*Order, error) {
	path := "order/new"
	params := map[string]interface{}{
		"symbol":   currencyPairToSymbol(pair),
		"amount":   amount,
		"price":    price,
		"side":     side,
		"type":     orderType,
		"exchange": "bitfinex"}

	var respmap map[string]interface{}
	req, err := s.client.newAuthenticatedRequest("POST", path.Join("orders", symbol), nil)

	err := doAuthenticatedRequest("POST", path, params, &respmap)
	if err != nil {
		return nil, err
	}

	order := new(Order)
	order.Currency = pair
	order.OrderID = ToInt(respmap["id"])
	order.Amount = ToFloat64(amount)
	order.Price = ToFloat64(price)
	order.AvgPrice = ToFloat64(respmap["avg_execution_price"])
	order.DealAmount = ToFloat64(respmap["executed_amount"])
	order.Status = ORDER_UNFINISH

	switch side {
	case "buy":
		if orderType == "limit" || orderType == "exchange limit" {
			order.Side = BUY
		} else {
			order.Side = BUY_MARKET
		}
	case "sell":
		if orderType == "limit" || orderType == "exchange limit" {
			order.Side = SELL
		} else {
			order.Side = SELL_MARKET
		}

	}
	return order, nil
}

func (bfx *Bitfinex) LimitBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("exchange limit", "buy", amount, price, currencyPair)
}

func (bfx *Bitfinex) LimitSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("exchange limit", "sell", amount, price, currencyPair)
}

func (bfx *Bitfinex) MarketBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("exchange market", "buy", amount, price, currencyPair)
}

func (bfx *Bitfinex) MarketSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
	return bfx.placeOrder("exchange market", "sell", amount, price, currencyPair)
}

func TestOrdersHistory(t *testing.T) {
	httpDo = func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
				[
					[4419360502,null,83283216761,"tIOTBTC",1508281683000,1508281731000,63938,63938,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.0000843,0,0,0,null,null,null,0,0,null],
					[4419354239,null,83265164211,"tIOTBTC",1508281665000,1508281674000,63976,63976,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.00008425,0,0,0,null,null,null,0,0,null],
					[4419339620,null,83217673277,"tIOTBTC",1508281618000,1508281653000,64014,64014,"EXCHANGE LIMIT",null,null,null,null,"CANCELED",null,null,0.0000842,0,0,0,null,null,null,0,0,null]
				]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	orders, err := NewClient().Orders.History(TradingPrefix + IOTBTC)

	if err != nil {
		t.Error(err)
	}

	if len(orders) != 3 {
		t.Errorf("expected three orders but got %d", len(orders))
	}

	_, err = NewClient().Orders.History("")
	if err == nil {
		t.Errorf("expected error when supplying empty symbol but got none")
	}
}
