package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Status int
	Body   string
}

func main() {
	lambda.Start(handler)
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = string('*')
	return headers
}

func handleRespond(s int, b string) (events.APIGatewayProxyResponse, error) {
	bytes, err := json.Marshal(b)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Headers:    getHeaders(),
		Body:       string(bytes),
		StatusCode: s,
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	path := &request.Path
	body := &request.Body

	if strings.Contains(*path, "hello") {
		r := hello()
		return handleRespond(r.Status, r.Body)
	} else if strings.Contains(*path, "hi") {
		r, err := hi(body)
		if err != nil {
			return handleRespond(http.StatusInternalServerError, err.Error())
		}
		return handleRespond(r.Status, r.Body)
	}

	return handleRespond(http.StatusNotFound, "")
}
