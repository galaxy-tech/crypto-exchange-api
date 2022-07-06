package ftx

import (
	"encoding/json"
	chttp "github.com/galaxy-tech/crypto-exchange-api/http"
	"github.com/pkg/errors"
)

type MarketData struct {
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	Last float64 `json:"last"`
	Ask float64 `json:"ask"`
	Bid float64 `json:"bid"`
	Price float64 `json:"price"`
	Type string `json:"type"`
}

type MarketDataResult struct {
	Success bool `json:"success"`
	Result MarketData `json:"result"`
}

func QueryMarketData(client *chttp.HTTPClient, market string) (*MarketData, error) {

	url := "https://ftx.com/api/markets/" + market

	if client == nil {
		client = chttp.NewHTTPClient()
	}
	dat, err := client.GetHTTP(url)
	if err != nil {
		return nil, errors.Wrap(err, "http client error")
	}

	mdr := &MarketDataResult{}
	err = json.Unmarshal(dat, &mdr)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal error")
	}

	if not(mdr.Success) {
		return nil, errors.New("FTX API returned failure")
	}

	return &mdr.Result, nil
}


type MarketDataAskPrice struct {
	Ask float64 `json:"ask"`
}
type MarketDataAskPriceResult struct {
	Success bool `json:"success"`
	Result MarketDataAskPrice `json:"result"`
}
func QueryAskPrice(client *chttp.HTTPClient, market string) (float64, error) {

	url := "https://ftx.com/api/markets/" + market

	if client == nil {
		client = chttp.NewHTTPClient()
	}
	dat, err := client.GetHTTP(url)
	if err != nil {
		return 0.0, errors.Wrap(err, "http client error")
	}

	md := &MarketDataAskPriceResult{}
	err = json.Unmarshal(dat, &md)
	if err != nil {
		return 0.0, errors.Wrap(err, "json unmarshal error")
	}

	if not(md.Success) {
		return 0.0, errors.New("FTX API returned failure")
	}

	return md.Result.Ask, nil
}

func not(flag bool) bool {
	return !flag
}
