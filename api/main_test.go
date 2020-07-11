package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestMain(t *testing.T) {
	response, err := handler(events.APIGatewayProxyRequest{
		Path:                  "v1/hello",
		HTTPMethod:            http.MethodGet,
		Headers:               map[string]string{},
		QueryStringParameters: map[string]string{},
		PathParameters:        map[string]string{},
	})
	if !(response.StatusCode == 200 && err == nil) {
		t.Error(err)
	}

	sampleBody := hiParameters{
		Name: "Maria",
	}

	bodyBytes, _ := json.Marshal(sampleBody)
	body := string(bodyBytes)

	response, err = handler(events.APIGatewayProxyRequest{
		Path:                  "v1/hi",
		HTTPMethod:            http.MethodPost,
		Headers:               map[string]string{},
		QueryStringParameters: map[string]string{},
		PathParameters:        map[string]string{},
		Body:                  body,
	})
	if !(response.StatusCode == 200 && err == nil) {
		t.Error(err)
	}

	response, err = handler(events.APIGatewayProxyRequest{
		Path:                  "v1/ciao",
		HTTPMethod:            http.MethodGet,
		Headers:               map[string]string{},
		QueryStringParameters: map[string]string{},
		PathParameters:        map[string]string{},
	})
	if !(response.StatusCode == 404 && err == nil) {
		fmt.Print(response)
		t.Error(err)
	}

}
