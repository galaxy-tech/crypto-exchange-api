package ftx

import (
	"encoding/json"
	"fmt"
	chttp "github.com/galaxy-tech/crypto-exchange-api/http"
	"github.com/pkg/errors"
	"time"
)

type CandleData struct {
	StartTime string	`json:"startTime"`
	Time float64	`json:"time"`
	Open float32	`json:"open"`
	High float32	`json:"high"`
	Low float32		`json:"low"`
	Close float32	`json:"close"`
	Volume float32	`json:"volume"`
}

type CandleDataResult struct {
	Success bool	`json:"success"`
	Result []CandleData	`json:"result"`
}

func Query15MinCandles(client *chttp.HTTPClient, market string, count uint64) ([]CandleData, error) {

	var resolution uint64 = 15*60
	endTime := uint64(time.Now().Unix())
	startTime := endTime - (count * resolution)
	url := fmt.Sprintf(
		"https://ftx.com/api/markets/%s/candles?resolution=%d&start_time=%d&end_time=%d",
		market, resolution, startTime, endTime)

	if client == nil {
		client = chttp.NewHTTPClient()
	}
	dat, err := client.GetHTTP(url)
	if err != nil {
		return nil, errors.Wrap(err, "http client error")
	}

	cdr := &CandleDataResult{}
	err = json.Unmarshal(dat, &cdr)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal error")
	}

	if not(cdr.Success) {
		return nil, errors.New("FTX API returned failure")
	}

	return cdr.Result[:count], nil
}
