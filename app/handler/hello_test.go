package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Hello))
	defer ts.Close()

	response, err := http.Get(ts.URL)
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Get unexpected error: %v\n", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code error: %v\n", response.StatusCode)
	}
	if fmt.Sprintf("%s", body) != "Hello World" {
		t.Errorf("Response body error: %v\n", response.StatusCode)
	}
}
