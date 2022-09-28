package api

import (
	"fmt"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type Track struct {
	ID       string
	Name     string
	Artists  map[string][]string
	Explicit bool
	Duration string
}

var artists = make(map[string][]string)

func GetPlaylistTracks(playlist *Playlist) (tracks []*Track) {
	fullPlaylist := describePlaylist(spotify.ID(playlist.ID))

	for _, track := range fullPlaylist.Tracks.Tracks {
		trackId := track.Track.ID
		trackName := track.Track.Name
		isExplicit := track.Track.Explicit
		simpleArtists := track.Track.Artists
		duration := track.Track.Duration

		for _, simpleArtist := range simpleArtists {
			artists[trackName] = append(artists[trackName], simpleArtist.Name)
		}

		tracks = append(tracks,
			&Track{
				ID:       string(trackId),
				Name:     trackName,
				Artists:  artists,
				Explicit: isExplicit,
				Duration: convertMillisecondsToMinutesAndSeconds(duration),
			})
	}

	return tracks
}

func ConvertTrackArtistListToSingleString(artists []string) string {
	return strings.Join(artists, ", ")
}

func convertMillisecondsToMinutesAndSeconds(milliseconds int) string {
	seconds := milliseconds / 1000
	minutes := seconds / 60
	return fmt.Sprint(minutes) + ":" + fmt.Sprint(seconds)[:len(fmt.Sprint(seconds))-1]
}
