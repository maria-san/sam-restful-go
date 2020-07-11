package main

import "net/http"

func hello() response {
	resp := response{
		status: http.StatusOK,
		body:   "Hello",
	}
	return resp
}
