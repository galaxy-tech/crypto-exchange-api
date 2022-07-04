package main

import (
	"fmt"
	"github.com/galaxy-tech/crypto-exchange-api/ftx"
)

func main() {
	fmt.Println(ftx.QueryMarketData(nil, "MATIC/USD"))
	fmt.Println(ftx.Query15MinCandles(nil, "MATIC/USD", 5))
}
