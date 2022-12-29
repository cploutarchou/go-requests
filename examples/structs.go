package examples

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
