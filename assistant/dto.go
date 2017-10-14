package assistant



type History struct {
	Order     string `json:"order"`
	OrderType string `json:"orderType"`
	Pair      string `json:"pair"`
	Price     string `json:"price"`
	VolumeRem string `json:"volumeRem"`
	CostRem   string `json:"costRem"`
	Status    string `json:"status"`
	Opened    string `json:"opened"`
}


type TradeReq struct {

}

type TradeRes struct {

}

type CurrenciesRes struct {

}