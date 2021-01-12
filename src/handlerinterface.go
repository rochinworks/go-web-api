package main

import (
	"net/http"
)

// Handler interface makes mocking out the http layer in tests easier
type Handler interface {
	BaseHandler() http.HandlerFunc
}
