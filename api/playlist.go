package api

import (
	"log"
	"strconv"

	"github.com/zmb3/spotify"
)

func DescribePlaylist(playlistID spotify.ID) *spotify.FullPlaylist {
	client := initClient()

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	return playlist
}

func PlayListFollowersCount(playlistID spotify.ID) string {
	playlist := DescribePlaylist(playlistID)
	playListFollowersCount := strconv.FormatUint(uint64(playlist.Followers.Count), 10)
	return playListFollowersCount
}

func PlaylistName(playlistID spotify.ID) string {
	playlist := DescribePlaylist(playlistID)
	return playlist.SimplePlaylist.Name
}
