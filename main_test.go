package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200")
	}
	defer resp.Body.Close()

	expected := "foo"

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected foo but got %s", string(b))
	}

}
