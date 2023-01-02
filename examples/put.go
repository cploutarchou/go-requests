package examples

import (
	"encoding/json"
)

func updatePet(item *Pet) (*updateRes, error) {
	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	res, err := jsonContentClient.Put(baseURL+"/pet", nil, data)
	if err != nil {
		return nil, err
	}

	var response updateRes
	err = res.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
