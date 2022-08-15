package ftx

import (
	"github.com/galaxy-tech/go-ftx/rest/private/orders"
	"github.com/galaxy-tech/go-ftx/types"
	"github.com/pkg/errors"
)

func PlaceLimitBuyOrder(rc *RestClient, market, clientID string, price, size float64) (*orders.ResponseForPlaceOrder, error) {
	dat, err := rc.client.PlaceOrder(&orders.RequestForPlaceOrder{
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
		return nil, errors.Wrap(err, "error placing buy limit order")
	}
	return dat, nil
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

func PlaceTrailingStopSellOrder(rc *RestClient, market string, size, trailValue float64 ) (*orders.ResponseForPlaceTriggerOrder, error) {
	if trailValue >= 0 {
		return nil, errors.New("TrailValue should be negative for Sell Trailing Stop order")
	}

	dat, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "trailingStop",
		Side:             "sell",
		Size:             size,
		TrailValue:       trailValue,
		ReduceOnly:       true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error placing sell trailing stop order")
	}
	return dat, nil
}

func PlaceTrailingStopBuyOrder(rc *RestClient, market string, size, trailValue float64 ) (*orders.ResponseForPlaceTriggerOrder, error) {
	if trailValue >= 0 {
		return nil, errors.New("TrailValue should be negative for Sell Trailing Stop order")
	}

	dat, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "trailingStop",
		Side:             "buy",
		Size:             size,
		TrailValue:       trailValue,
		ReduceOnly:       true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error placing sell trailing stop order")
	}
	return dat, nil
}

func PlaceStopLossOrder(rc *RestClient, market, action string, size, triggerPrice, orderPrice float64 ) (*orders.ResponseForPlaceTriggerOrder, error) {
	if action != "buy" && action != "sell" {
		return nil, errors.New("Action should be buy or sell")
	}

	dat, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "stop",
		Side:             action,
		Size:             size,
		TriggerPrice:     triggerPrice,
		OrderPrice: 	  orderPrice,
		ReduceOnly:       true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error placing sell trailing stop order")
	}
	return dat, nil
}

func PlaceEntryStopLossOrder(rc *RestClient, market, action string, size, triggerPrice, orderPrice float64 ) (*orders.ResponseForPlaceTriggerOrder, error) {
	if action != "buy" && action != "sell" {
		return nil, errors.New("Action should be buy or sell")
	}

	dat, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "stop",
		Side:             action,
		Size:             size,
		TriggerPrice:     triggerPrice,
		OrderPrice: 	  orderPrice,
		ReduceOnly:       false,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error placing sell trailing stop order")
	}
	return dat, nil
}

func PlaceTakeProfitOrder(rc *RestClient, market, action string, size, triggerPrice, orderPrice float64 ) (*orders.ResponseForPlaceTriggerOrder, error) {

	if action != "buy" && action != "sell" {
		return nil, errors.New("Action should be buy or sell")
	}

	dat, err := rc.client.PlaceTriggerOrder(&orders.RequestForPlaceTriggerOrder{
		Market:           market,
		Type:             "takeProfit",
		Side:             action,
		Size:             size,
		TriggerPrice:     triggerPrice,
		OrderPrice: 	  orderPrice,
		ReduceOnly:       true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error placing sell trailing stop order")
	}
	return dat, nil
}