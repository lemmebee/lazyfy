package api

import (
	"context"
	"log"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type Track struct {
	Name    string
	Artists map[string][]string
	// ExternalURLs map[string]string
	// PreviewURL   string
	// Duration     int
	Explicit bool
}

type Artist struct {
	Name string
}

type Playlist struct {
	ID   spotify.ID
	Name string
}

var (
	ctx                                 = context.Background()
	playlists     []Playlist            = nil
	playlistIDs   []string              = nil
	playlistNames []string              = nil
	tracks        []Track               = nil
	artists       []map[string][]string = nil
	artistNames   []string              = nil
)

func DescribePlaylist(playlist Playlist) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(ctx, playlist.ID)
	if err != nil {
		log.Fatalf("Error fetching full playlist: %v", err)
	}

	return fullPlaylist
}

// func GetPlayListFollowersCount(playlist Playlist) (playListFollowersCount string) {
// 	fullPlaylist := DescribePlaylist(playlist)
// 	playListFollowersCount = strconv.FormatUint(uint64(fullPlaylist.Followers.Count), 10)

// 	return playListFollowersCount
// }

// func GetPlayListTrackCount(playlist Playlist) (playlistTrackCount string) {
// 	fullPlaylist := DescribePlaylist(playlist)
// 	playlistTrackCount = strconv.FormatUint(uint64(fullPlaylist.Tracks.Total), 10)

// 	return playlistTrackCount
// }

// func GetSimplePlaylistsWithCountry(countryCode string) (message string, simplePlaylists []spotify.SimplePlaylist) {
// 	message, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx, spotify.Country(countryCode))
// 	if err != nil {
// 		log.Default().Fatalln("Error fetching simple playlist pages:", err)
// 	}
// 	simplePlaylists = simplePlaylistPages.Playlists
// 	return message, simplePlaylists
// }

func GetSimplePlaylists() (message string, simplePlaylists []spotify.SimplePlaylist) {
	message, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching simple playlist pages:", err)
	}
	simplePlaylists = simplePlaylistPages.Playlists
	return message, simplePlaylists
}

func GetPlaylists() (playlists []Playlist) {

	_, simplePlaylists := GetSimplePlaylists()

	for _, simplePlaylist := range simplePlaylists {
		playlists = append(playlists,
			Playlist{
				ID:   simplePlaylist.ID,
				Name: simplePlaylist.Name,
			})
	}

	return playlists
}

// func GetPlaylistIDs(playlists []Playlist) (playlistIDs []string) {
// 	for _, playlist := range playlists {
// 		playlistID := playlist.ID
// 		playlistIDs = append(playlistIDs, playlistID.String())
// 	}

// 	return playlistIDs
// }

// func GetPlaylistNames(playlists []Playlist) (playlistNames []string) {
// 	for _, playlist := range playlists {
// 		playlistName := playlist.Name
// 		playlistNames = append(playlistNames, playlistName)
// 	}
// 	return playlistNames
// }

func GetPlaylistTracks(playlist Playlist) (tracks []Track) {
	fullPlaylist := DescribePlaylist(playlist)

	for _, track := range fullPlaylist.Tracks.Tracks {
		trackName := track.Track.Name
		isExplicit := track.Track.Explicit
		simpleArtists := track.Track.Artists

		trackArtist := make(map[string][]string)
		for _, simpleArtist := range simpleArtists {
			trackArtist[trackName] = append(trackArtist[trackName], simpleArtist.Name)
		}

		artists = append(artists, trackArtist)

		tracks = append(tracks,
			Track{
				Name:     trackName,
				Artists:  trackArtist,
				Explicit: isExplicit,
			})
	}
	return tracks
}

func ConvertTrackArtistListToSingleString(artists []string) string {
	return strings.Join(artists, ", ")
}
