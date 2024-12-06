package handlers

import (
	"SpotigramAPI/internal/config"
	spotify "SpotigramAPI/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CurrentlyPlayingHandler(w http.ResponseWriter, r *http.Request, cfg config.Config) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var request config.Request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cant decode json"))
		return
	}
	response, statusCode := spotify.CurrentlyPlaying(request.Access_token, cfg)
	if statusCode != http.StatusOK {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(statusCode)
		return
	}

	log.Println(request)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Something went wrong: ", err)
	}

}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ping")
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Логируем метод запроса, URL и заголовки
	log.Printf("Method: %s, URL: %s, Remote Address: %s", r.Method, r.URL.Path, r.RemoteAddr)

	// Логируем заголовки запроса
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("Header: %s: %s", name, value)
		}
	}

}
