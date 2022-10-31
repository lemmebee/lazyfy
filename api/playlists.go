package api

import (
	"context"
	"log"
	"reflect"
	"strconv"

	"github.com/zmb3/spotify/v2"
)

type Playlist struct {
	ID    string
	Name  string
	Likes string
}

var Ctx = context.Background()

const description string = "lazyfy did this! Check it out! https://github.com/ehabshaaban/lazyfy"

func describePlaylist(playlistID spotify.ID) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(Ctx, playlistID)
	if err != nil {
		log.Fatalf("Error fetching full playlist: %v", err)
	}

	return fullPlaylist
}

func getSimplePlaylists() (simplePlaylists []spotify.SimplePlaylist) {
	_, simplePlaylistPages, err := Client.FeaturedPlaylists(Ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching simple playlist pages:", err)
	}
	simplePlaylists = simplePlaylistPages.Playlists

	return simplePlaylists
}

func GetPlaylists() (playlists []*Playlist) {
	simplePlaylists := getSimplePlaylists()

	for _, simplePlaylist := range simplePlaylists {

		if !reflect.DeepEqual(simplePlaylist, spotify.SimplePlaylist{}) {
			playlists = append(playlists,
				&Playlist{
					ID:    string(simplePlaylist.ID),
					Name:  simplePlaylist.Name,
					Likes: getPlayListLikes(simplePlaylist.ID),
				})
		}
	}

	return playlists
}

func getPlayListLikes(playlistID spotify.ID) string {
	fullPlaylist := describePlaylist(playlistID)
	likes := strconv.FormatUint(uint64(fullPlaylist.Followers.Count), 10)

	return likes
}

func CreatePlaylistForUser(playlistName string, isPlaylistPublic bool) *spotify.FullPlaylist {
	if playlistName != "" {
		fullPlaylist, err := Client.CreatePlaylistForUser(Ctx, getUserID(), playlistName, description, isPlaylistPublic, false)
		if err != nil {
			log.Fatal(err)
		}
		return fullPlaylist
	} else {
		fullPlaylist, err := Client.CreatePlaylistForUser(Ctx, getUserID(), "lazyfy me daddy ;)", description, isPlaylistPublic, false)
		if err != nil {
			log.Fatal(err)
		}
		return fullPlaylist
	}
}

func AddTracksToPlaylist(playlistID spotify.ID, selectedTracks map[string]string) {
	keys := make([]spotify.ID, 0, len(selectedTracks))
	for k := range selectedTracks {
		keys = append(keys, spotify.ID(k))
	}
	Client.AddTracksToPlaylist(Ctx, playlistID, keys...)
}
