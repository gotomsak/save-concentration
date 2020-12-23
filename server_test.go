package main

import (
	"log"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {

	e := router()

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	log.Print(req)
}
