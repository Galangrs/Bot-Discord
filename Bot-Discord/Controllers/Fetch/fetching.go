package fetching

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func SendHTTPRequest(method, url string, data []byte) ([]byte, error) {
	var requestBody *bytes.Reader
	if data != nil {
		requestBody = bytes.NewReader(data)
	}

	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check for HTTP error status
	if response.StatusCode >= 400 {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return responseBody, errors.New(string(responseBody))
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
