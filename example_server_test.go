package main

import (
	"net/http"
	"testing"
)

func TestExampleServer(t *testing.T) {
	http.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	t.Log("started server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		t.Fatal(err)
	}
}
