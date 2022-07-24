package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SpotifyID     string
	SpotifySecret string
	RedirectURI   string
}

// New returns a Config struct with env variables
func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Config{
		SpotifyID:     "",
		SpotifySecret: "",
		RedirectURI:   "",
	}
	if spotifyID, present := os.LookupEnv("SPOTIFY_ID"); present {
		config.SpotifyID = spotifyID
	}
	if spotifySecret, present := os.LookupEnv("SPOTIFY_SECRET"); present {
		config.SpotifySecret = spotifySecret
	}
	if redirectURI, present := os.LookupEnv("REDIRECT_URI"); present {
		config.RedirectURI = redirectURI
	}

	return &config
}
