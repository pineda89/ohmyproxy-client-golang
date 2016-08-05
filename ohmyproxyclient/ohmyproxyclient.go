package ohmyproxyclient

import (
	"sync"
	"ohmyproxy-client-golang/domain"
	"ohmyproxy-client-golang/connector"
	"ohmyproxy-client-golang/jsonparser"
)

var defaultendpoint string = "http://46.101.170.161:8080"
var endpoint string = defaultendpoint



var mu sync.Mutex

type ohmyproxyclient struct {
	Endpoint string
}

func InitializeEndpoint(newendpoint string) {
	endpoint = newendpoint
}

func GetProxiesFromWS() *domain.Response {
	responseString := connector.SendGetRequestToEndpoint(completeEndpoint(endpoint));
	responseObject := jsonparser.ParseStringToResponse(responseString)
	return responseObject
}

func completeEndpoint(innerEndpoint string) string {
	return innerEndpoint + "/findProxies"
}