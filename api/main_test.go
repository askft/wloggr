package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/askft/wloggr/api/route"
)

func TestRouter(t *testing.T) {
	r := route.SetupRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Hello! This is a public API for use with wloggr.com.\n"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := route.SetupRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/undefined")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}
