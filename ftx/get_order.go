package ftx

import "github.com/galaxy-tech/go-ftx/rest/private/orders"

func GetTriggerOrder(rc *RestClient, ID string) (*orders.ResponseForOrderTriggers, error) {
	req := &orders.RequestForOrderTriggers{CID: ID}
	return rc.client.OrderTriggers(req)
}
