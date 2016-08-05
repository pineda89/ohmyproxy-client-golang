package connector

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func SendGetRequestToEndpoint(endpoint string) string {
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(htmlData)
}

