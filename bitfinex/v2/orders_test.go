package v2

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"derbysoft.com/derbysoft-rpc-go/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
)


func TestOrderService_History(t *testing.T) {
	bitfinex := NewClient().Credentials("eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt", "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb")
	orders,err:=bitfinex.Orders.History("")
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
