package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	expectedResult := response{
		Status: 200,
		Body:   "Hello",
	}
	result := hello()

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
