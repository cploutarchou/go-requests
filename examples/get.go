package examples

func findPetsByTag(tag string) (PetsTags, error) {
	client.QueryParams().Set("tags", tag)
	resp, err := client.Get(baseURL+"/pet/findByTags", nil)
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
