package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sguessou/go-httpclient/gohttp"
)

var (
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

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
	getUrls()

}

func getUrls() {
	resp, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
