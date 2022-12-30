package examples

import "time"

type PetTag struct {
	PhotoUrls []string `json:"photoUrls"`
	Name      string   `json:"name"`
	ID        int64    `json:"id"`
	Category  struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"category"`
	Tags []struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"tags"`
	Status string `json:"status"`
}
type PetsTags []PetTag

type Order struct {
	PetId    int64     `json:"petId"`
	Quantity int       `json:"quantity"`
	Id       int64     `json:"id"`
	ShipDate time.Time `json:"shipDate"`
	Complete bool      `json:"complete"`
	Status   string    `json:"status"`
}
type Pet struct {
	Category struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"tags"`
	Status string `json:"status"`
	ID     int64  `json:"id"`
}

type updateRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
