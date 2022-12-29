package examples

func placePetOrder(item *Order) (*Order, error) {
	res, err := jsonContentClient.Post(baseURL+"/store/order", nil, item)
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
