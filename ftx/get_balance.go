package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/wallet"
	"github.com/pkg/errors"
)

func GetBalance(rc *RestClient, coin string) (float64, error) {
	req := &wallet.RequestForBalances{}
	balances, err := rc.client.Balances(req)
	if err != nil {
		return 0.0, errors.Wrap(err, "error getting FTX balance")
	}

	for _, balance := range *balances {
		if balance.Coin == coin {
			return balance.Total, nil
		}
	}
	return 0.0, nil
}