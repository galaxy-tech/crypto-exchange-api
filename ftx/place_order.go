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

func PlaceLimitSellOrder(rc *RestClient, market, clientID string, price, size float64) (*orders.ResponseForPlaceOrder, error) {
	dat, err := rc.client.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.LIMIT,
		Market: market,
		Side:   types.SELL,
		Price:  price,
		Size:   size,
		// Optionals
		ClientID:   clientID,
		Ioc:        false,
		ReduceOnly: false,
		PostOnly:   false,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error placing buy limit order")
	}
	return dat, nil
}


func PlaceMarketBuyOrder(rc *RestClient, market, clientID string, size float64) (*orders.ResponseForPlaceOrder, error) {
	dat, err := rc.client.PlaceOrder(&orders.RequestForPlaceOrder{
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
		return nil, errors.Wrap(err, "error placing buy limit order")
	}
	return dat, nil
}

func PlaceTrailingStopSellOrder(rc *RestClient, market string, size, trailValue float64 ) error {
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

func PlaceTrailingStopBuyOrder(rc *RestClient, market string, size, trailValue float64 ) error {
	if trailValue >= 0 {
		return errors.New("TrailValue should be negative for Sell Trailing Stop order")
	}

	_, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "trailingStop",
		Side:             "buy",
		Size:             size,
		TrailValue:       trailValue,
		ReduceOnly:       true,
	})

	if err != nil {
		return errors.Wrap(err, "error placing sell trailing stop order")
	}
	return nil
}

func PlaceStopLossOrder(rc *RestClient, market, action string, size, triggerPrice, orderPrice float64 ) error {
	if action != "buy" && action != "sell" {
		return errors.New("Action should be buy or sell")
	}

	_, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "stop",
		Side:             action,
		Size:             size,
		TriggerPrice:     triggerPrice,
		OrderPrice: 	  orderPrice,
		ReduceOnly:       true,
	})

	if err != nil {
		return errors.Wrap(err, "error placing sell trailing stop order")
	}
	return nil
}

func PlaceTakeProfitOrder(rc *RestClient, market, action string, size, triggerPrice, orderPrice float64 ) error {

	if action != "buy" && action != "sell" {
		return errors.New("Action should be buy or sell")
	}

	_, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "takeProfit",
		Side:             action,
		Size:             size,
		TriggerPrice:     triggerPrice,
		OrderPrice: 	  orderPrice,
		ReduceOnly:       true,
	})

	if err != nil {
		return errors.Wrap(err, "error placing sell trailing stop order")
	}
	return nil
}