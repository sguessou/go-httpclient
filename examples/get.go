package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl   string `json:"current_user_url"`
	AuthorizationUrl string `json:"authorization_url"`
	RepositoryUrl    string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	resp, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", resp.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", resp.Status()))
	fmt.Println(fmt.Sprintf("Body: %s", resp.String()))

	var endpoints Endpoints
	if err := resp.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))

	return &endpoints, nil
}
