package mcdbcgoclient

type TickerData struct {
	High float32 `json:"high"`
	Low  float32 `json:"low"`
	Vol  float32 `json:"vol"`
	Last float32 `json:"last"`
	Buy  float32 `json:"buy"`
	Sell float32 `json:"sell"`
	Date int     `json:"date"`
}
type Ticker struct {
	Ticker TickerData `json:"ticker"`
}

type Orderbook struct {
	Asks [][]float32 `json:"asks"`
	Bids [][]float32 `json:"bids"`
}

type TradesData struct {
	Date   int     `json:"date"`
	Price  float32 `json:"price"`
	Amount float32 `json:"amount"`
	Tid    int     `json:"tid"`
	Type   string  `json:"type"`
}

type Trades struct {
	Collection []TradesData
}
