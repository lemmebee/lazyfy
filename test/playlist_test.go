package test

import (
	"context"
	"log"
	"testing"

	"github.com/ehabshaaban/lazyfy/api"
)

func TestFeaturedPlaylists(t *testing.T) {
	_, playlists, err := api.Client.FeaturedPlaylists(context.Background())
	if err != nil {
		log.Default().Println(playlists)
		log.Default().Fatalln("Expected to fetch playlists from Spotify!", err)
	}
}
