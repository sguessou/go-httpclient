package gohttpclient

import "github.com/sguessou/go-httpclient/gohttp"

func basicExample() {
	client := gohttp.New()

	client.Get()

}
