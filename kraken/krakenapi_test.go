package krakenapi

import (
	"encoding/base64"
	"net/url"
	"reflect"
	"testing"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/log"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"fmt"
	"time"
)

var publicAPI = New("RNL8qrMdKy+wRwCCR7cm5xHN09Bsew3snZIN3aW3rlnLPvtHTkCKvS+u", "WP90951w5I9uFCLabh8x0SqaKaqeTCe+orIez89Io/68R8i9Xh5lnQeSOsXtlTpf4KJ+ryf8kRMFHyRzuBpfSg==")

func TestMac(t *testing.T) {
	a := []byte("12")
	log.Debug(a)
	b := base64.StdEncoding.EncodeToString(getHMacSha512([]byte("12"), []byte("23")))
	log.Debug(b)
	bs := getSha256(a)
	log.Debug(base64.StdEncoding.EncodeToString(bs))
}

//UKszwkonYhwkibihsUIcVI25nSrJwKo0loW7H2u/NPFmp4jGwsEjdVxNrHgDTNsQ4pdQJbrK+D3Hsq94/yyRYg==
//UKszwkonYhwkibihsUIcVI25nSrJwKo0loW7H2u/NPFmp4jGwsEjdVxNrHgDTNsQ4pdQJbrK+D3Hsq94/yyRYg==
func TestSignature(t *testing.T) {
	log.Debug(fmt.Sprintf("%d", time.Now().UnixNano()))
	values := url.Values{
		"nonce": {"123"},
	}
	urlPath:="1"
	sig:=createSignature(urlPath,values,[]byte("1"))
	log.Debug(sig)
}

func TestCreateSignature(t *testing.T) {
	expectedSig := "Uog0MyIKZmXZ4/VFOh0g1u2U+A0ohuK8oCh0HFUiHLE2Csm23CuPCDaPquh/hpnAg/pSQLeXyBELpJejgOftCQ=="
	urlPath := "/0/private/"
	secret, _ := base64.StdEncoding.DecodeString("SECRET")
	values := url.Values{
		"TestKey": {"TestValue"},
	}

	sig := createSignature(urlPath, values, secret)

	if sig != expectedSig {
		t.Errorf("Expected Signature to be %s, got: %s\n", expectedSig, sig)
	}
}

func TestTime(t *testing.T) {
	resp, err := publicAPI.Time()
	if err != nil {
		t.Errorf("Time() should not return an error, got %s", err)
	}

	if resp.Unixtime <= 0 {
		t.Errorf("Time() should return valid Unixtime, got %d", resp.Unixtime)
	}
}

func TestBalance(t *testing.T) {
	resp, err := publicAPI.Balance()
	if err != nil {
		t.Errorf("Time() should not return an error, got %s", err)
	}
	util.PrintDebugJson(resp)
	log.Debug(resp)
}

func TestAssets(t *testing.T) {
	_, err := publicAPI.Assets()
	if err != nil {
		t.Errorf("Assets() should not return an error, got %s", err)
	}
}

func TestAssetPairs(t *testing.T) {
	resp, err := publicAPI.AssetPairs()
	if err != nil {
		t.Errorf("AssetPairs() should not return an error, got %s", err)
	}

	if resp.XXBTZEUR.Base+resp.XXBTZEUR.Quote != XXBTZEUR {
		t.Errorf("AssetPairs() should return valid response, got %+v", resp.XXBTZEUR)
	}
}

func TestTicker(t *testing.T) {
	resp, err := publicAPI.Ticker(XETHZUSD)
	if err != nil {
		t.Errorf("Ticker() should not return an error, got %s", err)
	}

	util.PrintDebugJson(resp)
	if resp.XETHZUSD.OpeningPrice == 0 {
		t.Errorf("Ticker() should return valid OpeningPrice, got %+v", resp.XETHZUSD.OpeningPrice)
	}
}

func TestQueryTime(t *testing.T) {
	result, err := publicAPI.Query("Time", map[string]string{})
	resultKind := reflect.TypeOf(result).Kind()

	if err != nil {
		t.Errorf("Query should not return an error, got %s", err)
	}
	if resultKind != reflect.Map {
		t.Errorf("Query `Time` should return a Map, got: %s", resultKind)
	}
}

func TestQueryTicker(t *testing.T) {
	result, err := publicAPI.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})
	resultKind := reflect.TypeOf(result).Kind()

	if err != nil {
		t.Errorf("Query should not return an error, got %s", err)
	}

	if resultKind != reflect.Map {
		t.Errorf("Query `Ticker` should return a Map, got: %s", resultKind)
	}
}

func TestQueryTrades(t *testing.T) {
	result, err := publicAPI.Trades(XXBTZEUR, 1495777604391411290)

	if err != nil {
		t.Errorf("Trades should not return an error, got %s", err)
	}

	if result.Last == 0 {
		t.Errorf("Returned parameter `last` should always have a value...")
	}

	if len(result.Trades) > 0 {
		for _, trade := range result.Trades {
			if trade.Buy == trade.Sell {
				t.Errorf("Trade should be buy or sell")
			}
			if trade.Market == trade.Limit {
				t.Errorf("Trade type should be market or limit")
			}
		}
	}
}

func TestQueryDepth(t *testing.T) {
	pair := "XETHZEUR"
	count := 10
	result, err := publicAPI.Depth(pair, count)
	if err != nil {
		t.Errorf("Depth should not return an error, got %s", err)
	}

	resultType := reflect.TypeOf(result)

	if resultType != reflect.TypeOf(&OrderBook{}) {
		t.Errorf("Depth should return an OrderBook, got %s", resultType)
	}

	if len(result.Asks) > count {
		t.Errorf("Asks length must be less than count , got %s > %s", len(result.Asks), count)
	}

	if len(result.Bids) > count {
		t.Errorf("Bids length must be less than count , got %s > %s", len(result.Bids), count)
	}
}
