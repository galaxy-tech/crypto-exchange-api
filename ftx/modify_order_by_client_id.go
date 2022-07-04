package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/orders"
	"github.com/pkg/errors"
)

func ModifyLimitBuyOrder(rc *RestClient, market, clientID string, price, size float64) error {
	_, err := rc.client.ModifyOrder(&orders.RequestForModifyOrder{
		Price:  price,
		Size:   size,
		ClientID:   clientID,
	})
	if err != nil {
		return errors.Wrap(err, "error placing buy limit order")
	}
	return nil
}