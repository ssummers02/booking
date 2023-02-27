package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func newRequest(url string, jsonStr []byte, method string) (int, []byte) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return resp.StatusCode, body
}
func TestGetCities(t *testing.T) {
	// go app.Run()
	url := "http://localhost:8080/api/cities"
	code, body := newRequest(url, nil, http.MethodGet)
	checkResponseCode(t, http.StatusOK, code)
	if len(body) == 0 {
		t.Errorf("Expected non-empty response body")
	}
	t.Log(string(body))
}
