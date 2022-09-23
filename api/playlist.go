package api

import (
	"context"
	"log"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type Track struct {
	Name     string
	Artists  map[string][]string
	Explicit bool
}

type Playlist struct {
	ID   string
	Name string
}

var (
	ctx     = context.Background()
	artists = make(map[string][]string)
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

func GetPlaylistTracks(playlist *Playlist) (tracks []*Track) {
	fullPlaylist := DescribePlaylist(playlist)

	for _, track := range fullPlaylist.Tracks.Tracks {
		trackName := track.Track.Name
		isExplicit := track.Track.Explicit
		simpleArtists := track.Track.Artists

		for _, simpleArtist := range simpleArtists {
			artists[trackName] = append(artists[trackName], simpleArtist.Name)
		}

		tracks = append(tracks,
			&Track{
				Name:     trackName,
				Artists:  artists,
				Explicit: isExplicit,
			})
	}

	return tracks
}

func ConvertTrackArtistListToSingleString(artists []string) string {
	return strings.Join(artists, ", ")
}
