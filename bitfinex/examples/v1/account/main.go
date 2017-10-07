package main

// Set BFX_APIKEY and BFX_SECRET as :
//
// export BFX_API_KEY=YOUR_API_KEY
// export BFX_API_SECRET=YOUR_API_SECRET
//
// you can obtain it from https://www.bitfinex.com/api

import (
	"fmt"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func main() {
	//key := os.Getenv("BFX_API_KEY")
	//secret := os.Getenv("BFX_API_SECRET")
	key:="eRNLi8wYH0SXncyDUkfmWs99CVvjqSUnQ6KcBqnwhDt"
	secret:="4qaFdWxqA6whubyPI92DTWSnDKbiABuMP7CnixmM2Vb"
	//secret:="2"
	client := bitfinex.NewClient().Auth(key, secret)
	info, err := client.Wallet.All()
	//payload:=client.SignPayload("1")
	//log.Debug(payload)
	util.PrintDebugJson(info)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}
}
