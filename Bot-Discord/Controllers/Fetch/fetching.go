package fetching

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func SendHTTPRequest(method, url string, data []byte, headers map[string]string) ([]byte, error) {
	var requestBody *bytes.Reader
	if data != nil {
		requestBody = bytes.NewReader(data)
	} else {
		requestBody = bytes.NewReader([]byte{}) // or bytes.NewReader(nil)
	}
	

	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	request.Header.Add("Content-Type", "application/json")

	// Add custom headers
	for key, value := range headers {
		request.Header.Add(key, value)
	}

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
		errMsg := string(responseBody)
		return responseBody, errors.New("HTTP error: " + errMsg)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
