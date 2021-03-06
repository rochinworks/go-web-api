package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"go-web-api/src/handler"
)

func main() {

	// ========================================
	// start server with middleware
	// ========================================
	h := handler.NewHandler()
	r := httpRouter(h)

	server := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}
	log.Info("server up and running at ", server.Addr)
	log.Fatal(server.ListenAndServe())

}
