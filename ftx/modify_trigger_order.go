package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/orders"
	"github.com/pkg/errors"
)

func ModifyStopLossTriggerOrder(rc *RestClient, OrderID string, triggerPrice, orderPrice, size float64) error {
	_, err := rc.client.ModifyTriggerOrder(&orders.RequestForModifyTriggerOrder{
		OrderID:      OrderID,
		TriggerPrice: triggerPrice,
		OrderPrice:   orderPrice,
		Size:         size,
	})
	if err != nil {
		return errors.Wrap(err, "error placing buy limit order")
	}
	return nil
}
