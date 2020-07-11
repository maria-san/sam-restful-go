package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestHi(t *testing.T) {
	expectedResult := response{
		Status: 200,
		Body:   "Hi Maria!",
	}
	sampleBody := hiParameters{
		Name: "Maria",
	}
	bodyBytes, _ := json.Marshal(sampleBody)
	body := string(bodyBytes)

	result, err := hi(&body)
	if err != nil {
		t.Error(err)
	}

	bytesExpected, err := json.Marshal(expectedResult)
	if err != nil {
		t.Error(err)
	}
	bytesResult, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(bytesExpected, bytesResult) {
		t.Errorf("Result is not equal to the expected sample")
	}
}

func TestHiErr(t *testing.T) {
	sampleBody := "wrong body"

	bodyBytes, _ := json.Marshal(sampleBody)
	body := string(bodyBytes)

	_, err := hi(&body)
	if err == nil {
		t.Error(err)
	}

}
