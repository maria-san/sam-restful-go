package main

import "net/http"

func hello() response {
	resp := response{
		Status: http.StatusOK,
		Body:   "Hello",
	}
	return resp
}
