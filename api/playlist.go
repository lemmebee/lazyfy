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
	tracks []Track
}

var (
	ctx                        = context.Background()
	fPlaylists      []Playlist = nil
	fPlaylistTracks []Track    = nil
)

func DescribePlaylist(playlist Playlist) (fullPlaylist *spotify.FullPlaylist) {
	fullPlaylist, err := Client.GetPlaylist(ctx, playlist.id)
	if err != nil {
		log.Fatalf("Error fetching playlist: %v", err)
	}

	return fullPlaylist
}

func PlayListFollowersCount(playlist Playlist) (playListFollowersCount string) {
	fullPlaylist := DescribePlaylist(playlist)
	playListFollowersCount = strconv.FormatUint(uint64(fullPlaylist.Followers.Count), 10)

	return playListFollowersCount
}

func PlaylistName(playlist Playlist) (playlistName string) {
	fullPlaylist := DescribePlaylist(playlist)
	playlistName = fullPlaylist.SimplePlaylist.Name

	return playlistName
}

func GetFeaturedPlaylistsWithCountry(countryCode string) (message string, fPlaylists []spotify.SimplePlaylist) {
	message, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx, spotify.Country(countryCode))
	if err != nil {
		log.Default().Fatalln("Error fetching featured playlists", err)
	}
	fPlaylists = simplePlaylistPages.Playlists

	return message, fPlaylists
}

func GetFeaturedPlaylists() (message string, fPlaylists []spotify.SimplePlaylist) {
	message, simplePlaylistPages, err := Client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching featured playlists", err)
	}
	fPlaylists = simplePlaylistPages.Playlists

	return message, fPlaylists
}

func GetFeaturedPlaylist() (fPlaylists []Playlist) {

	_, fSimplePlaylists := GetFeaturedPlaylists()

	for _, fSimplePlaylist := range fSimplePlaylists {
		fPlaylists = append(fPlaylists, Playlist{id: fSimplePlaylist.ID})
	}
	for _, fPlaylist := range fPlaylists {
		fFullPlaylist := DescribePlaylist(fPlaylist)

		for _, fPlaylistTrack := range fFullPlaylist.Tracks.Tracks {
			fTrackName := fPlaylistTrack.Track.Name
			fPlaylistTracks = append(fPlaylistTracks, Track{name: fTrackName})
		}
	}
	fPlaylists = append(fPlaylists, Playlist{tracks: fPlaylistTracks})

	return fPlaylists
}
