package api

import (
	"context"
	"log"
	"strconv"

	"github.com/zmb3/spotify/v2"
)

type Track struct {
	name         string
	artists      []Artist
	externalURLs map[string]string
	previewURL   string
	duration     int
	explicit     bool
}

type Artist struct {
	name string
}

type Playlist struct {
	id     spotify.ID
	name   string
	tracks []Track
}

var (
	ctx                       = context.Background()
	playlists      []Playlist = nil
	playlistIDs    []string   = nil
	playlistNames  []string   = nil
	playlistTracks []Track    = nil
)

func DescribePlaylist(playlist Playlist) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(ctx, playlist.id)
	if err != nil {
		log.Fatalf("Error fetching full playlist: %v", err)
	}

	return fullPlaylist
}

func GetPlayListFollowersCount(playlist Playlist) (playListFollowersCount string) {
	fullPlaylist := DescribePlaylist(playlist)
	playListFollowersCount = strconv.FormatUint(uint64(fullPlaylist.Followers.Count), 10)

	return playListFollowersCount
}

func GetSimplePlaylistsWithCountry(countryCode string) (message string, simplePlaylists []spotify.SimplePlaylist) {
	message, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx, spotify.Country(countryCode))
	if err != nil {
		log.Default().Fatalln("Error fetching simple playlist pages:", err)
	}
	simplePlaylists = simplePlaylistPages.Playlists
	return message, simplePlaylists
}

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
				id:   simplePlaylist.ID,
				name: simplePlaylist.Name,
			})
	}
	for _, playlist := range playlists {
		fullPlaylist := DescribePlaylist(playlist)

		for _, track := range fullPlaylist.Tracks.Tracks {
			trackName := track.Track.Name
			playlistTracks = append(playlistTracks,
				Track{
					name: trackName,
				})
		}
	}
	playlists = append(playlists,
		Playlist{
			tracks: playlistTracks,
		})

	return playlists
}

func GetPlaylistIDs(playlists []Playlist) (playlistIDs []string) {
	for _, playlist := range playlists {
		playlistID := playlist.id
		playlistIDs = append(playlistIDs, playlistID.String())
	}
	return playlistIDs
}

func GetPlaylistNames(playlists []Playlist) (playlistNames []string) {
	for _, playlist := range playlists {
		playlistName := playlist.name
		playlistNames = append(playlistNames, playlistName)
	}
	return playlistNames
}
