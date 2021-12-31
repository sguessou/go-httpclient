package main

import (
	"fmt"

	"github.com/sguessou/go-httpclient/gohttp"
)

var (
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.
		NewBuilder().
		Build()

	// client.SetMaxIdleconnections(20)
	// client.SetConnectionTimeout(2 * time.Second)
	// client.SetResponseTimeout(4 * time.Millisecond)

	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Bearer abc-124")
	// client.SetHeaders(commonHeaders)

	return client
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	getUrls()
}

func getUrls() {
	resp, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status())
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())
}
