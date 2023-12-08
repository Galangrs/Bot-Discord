package fetching

import (
	"bytes"
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

	// header in request
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
