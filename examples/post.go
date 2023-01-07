package examples

import "encoding/json"

func placePetOrder(item Order) (*Order, error) {
	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	res, err := jsonContentClient.Post(baseURL+"/store/order", data)
	if err != nil {
		return nil, err
	}
	var order Order
	err = res.Unmarshal(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
