package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sguessou/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()

	resp, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

}
