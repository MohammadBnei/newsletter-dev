package main

import (
	"log"
	"net/http"

	"handler/function/infra"
	"handler/function/infra/config"
)

func main() {
	cfg := config.GetConfig()
	cfg.Local = true

	r, closeMongo := infra.Init()

	defer closeMongo()

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
