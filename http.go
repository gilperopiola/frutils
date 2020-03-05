package frutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendHTTPRequest(method, url, body string) (status int, response string) {
	request, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return resp.StatusCode, err.Error()
	}
	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return resp.StatusCode, string(b)
	}

	return 400, err.Error()
}

func SendHTTPRequestWithToken(method, url, body, token string) (status int, response string) {
	request, _ := http.NewRequest(method, url, bytes.NewReader([]byte(body)))

	request.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return resp.StatusCode, err.Error()
	}
	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return resp.StatusCode, string(b)
	}

	return 400, err.Error()
}
