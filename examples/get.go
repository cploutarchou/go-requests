package examples

func findPetsByTagJSON(tag string) (PetsTags, error) {
	jsonContentClient.QueryParams().Set("tags", tag)
	resp, err := jsonContentClient.Get(baseURL+"/pet/findByTags", nil)
	if err != nil {
		return nil, err
	}
	var pets PetsTags
	err = resp.Unmarshal(&pets)
	if err != nil {
		return nil, err
	}
	return pets, nil
}
