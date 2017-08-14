package mcdbcgoclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var reqUrl = "https://www.mercadobitcoin.net/api/"

func GetTicker() (*Ticker, error) {
	resp, err := http.Get(reqUrl + "ticker")
	if err != nil {
		return nil, err
	}
	ticker := new(Ticker)
	err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		return nil, err
	}
	return ticker, nil
}

func GetOrderBook() (*Orderbook, error) {
	resp, err := http.Get(reqUrl + "orderbook")
	if err != nil {
		return nil, err
	}
	orderBook := new(Orderbook)
	err = json.NewDecoder(resp.Body).Decode(&orderBook)
	if err != nil {
		return nil, err
	}
	return orderBook, nil
}

func GetTrades() ([]TradesData, error) {
	resp, err := http.Get(reqUrl + "trades")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	trades := make([]TradesData, 0)
	err = json.Unmarshal(body, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}
