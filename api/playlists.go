package api

import (
	"context"
	"log"
	"strconv"

	"github.com/zmb3/spotify/v2"
)

type Playlist struct {
	ID    string
	Name  string
	Likes string
}

var ctx = context.Background()

func describePlaylist(playlistID spotify.ID) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(ctx, playlistID)
	if err != nil {
		log.Fatalf("Error fetching full playlist: %v", err)
	}

	return fullPlaylist
}

func getSimplePlaylists() (simplePlaylists []spotify.SimplePlaylist) {
	_, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching simple playlist pages:", err)
	}
	simplePlaylists = simplePlaylistPages.Playlists

	return simplePlaylists
}

func GetPlaylists() (playlists []*Playlist) {
	simplePlaylists := getSimplePlaylists()

	for _, simplePlaylist := range simplePlaylists {
		playlists = append(playlists,
			&Playlist{
				ID:    string(simplePlaylist.ID),
				Name:  simplePlaylist.Name,
				Likes: getPlayListLikes(simplePlaylist.ID),
			})
	}

	emptyPlaylist := &Playlist{}

	for i, p := range playlists {
		if p == emptyPlaylist {
			playlists = append(playlists[:i], playlists[i+1:]...)
		}
	}

	return playlists
}

func getPlayListLikes(playlistID spotify.ID) string {
	fullPlaylist := describePlaylist(playlistID)
	likes := strconv.FormatUint(uint64(fullPlaylist.Followers.Count), 10)

	return likes
}
