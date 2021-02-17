package utils

import (
	"encoding/json"
	"net/http"

	"github.com/leogsouza/bitcoin-telegram-bot/model"
)

// GetAPICall calls the Bitcoin price API and returns the result
func GetAPICall() (*model.Price, error) {
	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")

	p := &model.Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
