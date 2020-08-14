package api

import (
	"io/ioutil"
	"net/http"
)

// Generic function to trigger GET Requests
func GetData(url string, headers map[string]string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	var client = http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	return body, err
}
