package ftx

import "github.com/galaxy-tech/go-ftx/rest/private/orders"

func GetOpenOrders(rc *RestClient, market string) (*orders.ResponseForOpenOrder, error) {
	req := &orders.RequestForOpenOrder{
		ProductCode: market,
	}
	return rc.client.OpenOrder(req)
}

func GetOpenTriggerOrders(rc *RestClient, market, typ string) (*orders.ResponseForOpenTriggerOrders, error) {
	req := &orders.RequestForOpenTriggerOrders{
		ProductCode: market,
		Type:        typ,
	}
	return rc.client.OpenTriggerOrders(req)
}