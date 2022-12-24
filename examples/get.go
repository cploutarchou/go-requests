package examples

type Endpoints struct {
	CurrentUserURL                   string `json:"current_user_url"`
	AuthorizationsURL                string `json:"authorizations_url"`
	CurrentUserAuthorizationsHTMLURL string `json:"current_user_authorizations_html_url"`
	RepositoryURL                    string `json:"repository_url"`
	CodeSearchURL                    string `json:"code_search_url"`
	CommitSearchURL                  string `json:"commit_search_url"`
}

func GetEndpoint(url string) (*Endpoints, error) {
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var endpoints Endpoints
	err = res.Unmarshal(&endpoints)
	if err != nil {
		return nil, err
	}
	return &endpoints, nil
}
