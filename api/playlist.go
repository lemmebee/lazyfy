package api

import (
	"context"
	"log"
	"strconv"

	"github.com/zmb3/spotify/v2"
)

func DescribePlaylist(playlistID spotify.ID) *spotify.FullPlaylist {
	playlist, err := Client.GetPlaylist(context.Background(), playlistID)
	if err != nil {
		log.Fatalf("Error retrieving playlist data: %v", err)
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
