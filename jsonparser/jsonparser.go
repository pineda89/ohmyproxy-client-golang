package jsonparser

import (
	"ohmyproxy-client-golang/domain"
	"encoding/json"
)

func ParseStringToResponse(jsonString string) *domain.Response {
	response := new(domain.Response)
	jsonBytes := []byte(jsonString)
	json.Unmarshal(jsonBytes, response)
	return response
}