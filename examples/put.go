package examples

import "fmt"

func updatePet(item *Pet) (*updateRes, error) {
	res, err := jsonContentClient.Put(baseURL+"/pet", nil, &item)
	fmt.Println("kokos")
	fmt.Println(res.Status())
	if err != nil {
		return nil, err
	}

	fmt.Println(res.String())
	var response updateRes
	err = res.Unmarshal(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
