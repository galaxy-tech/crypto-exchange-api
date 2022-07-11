package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/orders"
	"github.com/pkg/errors"
)

type CancelResponse struct {
	Success bool `json:"success"`
	Result  string `json:"result"`
}

func CancelTriggerOrder(rc *RestClient, id string) error {

	req := &orders.RequestForCancelByID{
		TriggerOrderID: id,
	}
	dat, err := rc.client.CancelByID(req)
	if err != nil {
		return errors.Wrap(err, "error in cancel trigger order")
	}


	if *dat != "Order cancelled" {
		return errors.New("error cancelling trigger failed")
	}

	return nil
}

func CancelOrder(rc *RestClient, id int) error {

	req := &orders.RequestForCancelByID{
		OrderID: id,
	}
	dat, err := rc.client.CancelByID(req)
	if err != nil {
		return errors.Wrap(err, "error in cancel  order")
	}

	if *dat != "Order cancelled" {
		return errors.New("error cancelling trigger failed")
	}

	return nil
}