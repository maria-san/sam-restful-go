package main

import (
	"encoding/json"
	"net/http"
)

type hiParameters struct {
	Name string `json:"name"`
}

func hi(b *string) (response, error) {

	resp := response{}
	h := hiParameters{}
	err := json.Unmarshal([]byte(*b), &h)
	if err != nil {
		return resp, err
	}

	resp = response{
		Status: http.StatusOK,
		Body:   "Hi " + h.Name + "!",
	}
	return resp, nil
}
