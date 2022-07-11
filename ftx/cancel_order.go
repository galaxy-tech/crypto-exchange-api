package ftx

import (
	"encoding/json"
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

	resp := CancelResponse{}
	err = json.Unmarshal([]byte(*dat), &resp)
	if err != nil {
		return errors.New("error unmarshalling json response")
	}

	if resp.Success != true {
		return errors.New("error cancelling trigger failed")
	}

	return nil
}