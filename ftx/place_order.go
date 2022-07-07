package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/orders"
	"github.com/galaxy-tech/go-ftx/types"
	"github.com/pkg/errors"
)

func PlaceLimitBuyOrder(rc *RestClient, market, clientID string, price, size float64) error {
	_, err := rc.client.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.LIMIT,
		Market: market,
		Side:   types.BUY,
		Price:  price,
		Size:   size,
		// Optionals
		ClientID:   clientID,
		Ioc:        false,
		ReduceOnly: false,
		PostOnly:   false,
	})
	if err != nil {
		return errors.Wrap(err, "error placing buy limit order")
	}
	return nil
}

func PlaceMarketBuyOrder(rc *RestClient, market, clientID string, price, size float64) error {
	_, err := rc.client.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.MARKET,
		Market: market,
		Side:   types.BUY,
		Size:   size,
		// Optionals
		ClientID:   clientID,
		Ioc:        false,
		ReduceOnly: false,
		PostOnly:   false,
	})
	if err != nil {
		return errors.Wrap(err, "error placing buy limit order")
	}
	return nil
}

func PlaceTrailingStopSellOrder(rc *RestClient, market, clientID string, size, trailValue float64 ) error {
	if trailValue >= 0 {
		return errors.New("TrailValue should be negative for Sell Trailing Stop order")
	}

	_, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "trailingStop",
		Side:             "sell",
		Size:             size,
		TrailValue:       trailValue,
		ReduceOnly:       true,
	})

	if err != nil {
		return errors.Wrap(err, "error placing sell trailing stop order")
	}
	return nil
}