package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func MakeHTTPRequest(method, url string, headers map[string]string, body string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(responseBody), nil
}

// Similar functions for HTTPS, WebSocket, FTP...
