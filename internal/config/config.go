package config

import "os"

type Config struct {
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyEndpoint     string
}

type Request struct {
	Access_token string `json:"access_token"`
}

type Response struct {
	Error   error `json:"error"`
	Playing bool  `json:"is_playing"`
	Item    struct {
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Name string `json:"name"`
	} `json:"item"`
}

// no need for docker
// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Please create .env file")
// 	}
// }

func NewConfig() Config {
	return Config{
		SpotifyClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyEndpoint:     "https://api.spotify.com/v1/me",
	}
	// no need for docker
	// return Config{
	// 	SpotifyClientId:     getEnv("SPOTIFY_CLIENT_ID"),
	// 	SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET"),
	// 	SpotifyEndpoint:     "https://api.spotify.com/v1/me",
	// }
}

// no need for docker
// func getEnv(key string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	} else {
// 		log.Println("Missing: ", value)
// 		return ""
// 	}
// }
