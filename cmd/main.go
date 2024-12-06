package main

import (
	"SpotigramAPI/internal/config"
	"SpotigramAPI/internal/routes"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	log.Println(cfg.SpotifyClientId)
	log.Println(cfg.SpotifyClientSecret)
	routes.InitRoutes(cfg)

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
