package routes

import (
	"SpotigramAPI/internal/config"
	"SpotigramAPI/internal/handlers"
	"net/http"
)

func InitRoutes(cfg config.Config) {
	http.HandleFunc("/", handlers.RequestHandler)
	http.HandleFunc("/currently-playing", func(w http.ResponseWriter, r *http.Request) { handlers.CurrentlyPlayingHandler(w, r, cfg) })
	http.HandleFunc("/pong", handlers.Pong)
	http.HandleFunc("/ping", handlers.Ping)
}
