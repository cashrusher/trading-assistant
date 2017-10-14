package assistant

import (
	"testing"
	"github.com/kraken-go-api-client"
	"encoding/json"
	"derbysoft.com/derbysoft-rpc-go/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"github.com/cashrusher/trading-assistant/bitfinex"
)

func TestTranslate2KrakenHistory(t *testing.T) {
	o := &krakenapi.Order{}
	str := `{
       "refid": "",
       "userref": "",
       "status": "closed",
       "opentm": 1504954174.0996,
       "starttm": 0,
       "expiretm": 0,
       "descr": {
         "pair": "ETHUSD",
         "close": "",
         "leverage": "none",
         "order": "sell 0.02000000 ETHUSD @ market",
         "ordertype": "market",
         "price": "0",
         "price2": "0",
         "type": "sell"
       },
       "vol1": "",
       "vol_exec": "0.02",
       "cost": "6.14",
       "fee": "0.01",
       "price": "307.29",
       "stopprice.string": 0,
       "limitprice": "0",
       "misc": "",
       "oflags": "fciq",
       "closetm": 1504954175.205,
       "reason": ""
     }`
	if err := json.Unmarshal([]byte(str), o); err != nil {
		log.Error(err)
	}
	util.PrintDebugJson(o)
	h := getKrakenHistory("12344", *o)
	util.PrintDebugJson(h)
}

func TestTranslate2BitfinexHistory(t *testing.T) {
	o := &bitfinex.Order{}
	str := `{
  "id":448411365,
  "symbol":"btcusd",
  "exchange":"bitfinex",
  "price":"0.02",
  "avg_execution_price":"0.0",
  "side":"buy",
  "type":"exchange limit",
  "timestamp":"1444276597.0",
  "is_live":false,
  "is_cancelled":true,
  "is_hidden":false,
  "was_forced":false,
  "original_amount":"0.02",
  "remaining_amount":"0.02",
  "executed_amount":"0.0"
}`
	if err := json.Unmarshal([]byte(str), o); err != nil {
		log.Error(err)
	}
	util.PrintDebugJson(o)
	h := getBitfinexHistory(*o)
	util.PrintDebugJson(h)
}
