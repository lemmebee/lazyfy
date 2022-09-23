package api

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
)

type Playlist struct {
	ID   string
	Name string
}

var (
	ctx = context.Background()
)

func DescribePlaylist(playlist *Playlist) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(ctx, spotify.ID(playlist.ID))
	if err != nil {
		log.Fatalf("Error fetching full playlist: %v", err)
	}

	return fullPlaylist
}

func GetSimplePlaylists() (simplePlaylists []spotify.SimplePlaylist) {
	_, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching simple playlist pages:", err)
	}
	simplePlaylists = simplePlaylistPages.Playlists

	return simplePlaylists
}

func GetPlaylists() (playlists []*Playlist) {
	simplePlaylists := GetSimplePlaylists()

	for _, simplePlaylist := range simplePlaylists {
		playlists = append(playlists,
			&Playlist{
				ID:   string(simplePlaylist.ID),
				Name: simplePlaylist.Name,
			})
	}

	return playlists
}
