package api

import "strings"

type Track struct {
	ID       string
	Name     string
	Artists  map[string][]string
	Explicit bool
}

var artists = make(map[string][]string)

func GetPlaylistTracks(playlist *Playlist) (tracks []*Track) {
	fullPlaylist := describePlaylist(playlist)

	for _, track := range fullPlaylist.Tracks.Tracks {
		trackId := track.Track.ID
		trackName := track.Track.Name
		isExplicit := track.Track.Explicit
		simpleArtists := track.Track.Artists

		for _, simpleArtist := range simpleArtists {
			artists[trackName] = append(artists[trackName], simpleArtist.Name)
		}

		tracks = append(tracks,
			&Track{
				ID:       string(trackId),
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
