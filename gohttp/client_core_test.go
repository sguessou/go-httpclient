package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid request id received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid request id received")
	}
}

func TestGetRequesBody(t *testing.T) {
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshalling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json obtained")
		}

	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"hello", "world"}
		body, err := client.getRequestBody("application/xml", requestBody)

		res := string(body)
		exp := `<string>hello</string><string>world</string>`

		if err != nil {
			t.Error("no error expected when marshalling slice as xml")
		}

		if res != exp {
			t.Errorf("invalid xml obtained: expected %q, got %q instead", exp, res)
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := map[string]string{"foo": "bar", "hello": "world"}
		body, err := client.getRequestBody("", requestBody)

		res := string(body)
		exp := `{"foo":"bar","hello":"world"}`

		if err != nil {
			t.Error("no error expected when marshalling slice as xml")
		}

		if res != exp {
			t.Errorf("invalid json obtained: expected %q, got %q instead", exp, res)
		}
	})

}
