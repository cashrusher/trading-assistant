package assistant

//OT5DME	sell/market	ETH/USD	$307.29	0.02000000	$6.14	Closed	09-09-17 18:49:35 +0800
/*
`
"OT5DME-3D5QU-JT7X3V": {
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
     }
`
*/

type History struct {
	Time      string  `json:"time"` // Beijing time format is yyyy-MM-DD hh:mm:ss
	OrderID   string  `json:"orderID"`
	Platform  string  `json:"platform"`
	OrderType string  `json:"orderType"` // Sell or buy
	Pair      string  `json:"pair"`
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	Amount    float64 `json:"amount"`
	Fee       float64 `json:"fee"`
	Status    string  `json:"status"`
}

/**
{
    "platform":"kraken", //can be 'kraken' or 'bitfinex'
    "currency":"BTH", // it's the currency returned by the currencies api
    "amount":34.9982,
    "price":887.000293
}
*/
type TradeReq struct {
	Platform string  `json:"platform"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Price    float64 `json:"price"`
}

/**
{
  "status":"success",
  "info":{
    "tradeID":"123455djsdkh33"
  }
}
*/
type TradeRes struct {
	Status string `json:"status"`
	Info struct {
		TradeID string `json:"tradeID"`
	} `json:"info"`
}


type CurrenciesRes []string
