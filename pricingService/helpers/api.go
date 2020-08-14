package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}
