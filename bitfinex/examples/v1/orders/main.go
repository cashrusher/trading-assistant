package main

import (
	"fmt"
	"github.com/cashrusher/trading-assistant/bitfinex"
	"derbysoft.com/derbysoft-rpc-go/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
)

// Set BFX_APIKEY and BFX_SECRET as :
//
// export BFX_API_KEY=YOUR_API_KEY
// export BFX_API_SECRET=YOUR_API_SECRET
//
// you can obtain it from https://www.bitfinex.com/api

// WARNING: IF YOU RUN THIS EXAMPLE WITH A VALID KEY ON PRODUCTION
//          IT WILL SUBMIT AN ORDER !
//API KEY:eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt
//Api secret: 4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb

func main() {
	key := "eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt"
	secret := "4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb"
	client := bitfinex.NewClient().Auth(key, secret)

	// Sell 0.01BTC at $12.000
	data, err := client.Orders.Create(bitfinex.ETHUSD, -0.001, 12000, bitfinex.OrderTypeMarket)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", data)
	}

	b,err:=client.Balances.All()
	if err!=nil{
		log.Error(err)
	}
	util.PrintDebugJson(b)
}
