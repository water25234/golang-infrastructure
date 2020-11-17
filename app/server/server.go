package coreserver

import (
	"log"
	"net/http"

	// "net/http"

	// "github.com/apex/gateway"

	"github.com/water25234/golang-infrastructure/config"
	"github.com/water25234/golang-infrastructure/router"
)

func init() {
	config.SetAppConfig()
}

// StartServer mean start server
func StartServer() {
	log.Fatal(http.ListenAndServe(
		config.GetAppConfig().GoAddrPort,
		router.SetupRouter(),
	))
}
