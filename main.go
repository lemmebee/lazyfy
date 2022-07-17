package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
	"github.com/ehabshaaban/lazyfy/ui"
	"github.com/zmb3/spotify"
)

const (
	playlistID spotify.ID = "4UscizPrS9cosaO4a4mbiF"
)

func main() {

	playlist := api.DescribePlaylist(playlistID)

	tracks := []list.Item{}

	for _, track := range playlist.Tracks.Tracks {

		for _, artist := range track.Track.SimpleTrack.Artists {

			tracks = append(tracks, ui.Track{Name: track.Track.SimpleTrack.Name, Artist: artist.Name})
		}
	}

	m := ui.Model{List: list.New(tracks, list.NewDefaultDelegate(), 0, 0)}

	m.List.Title = "You are seeing " + api.PlaylistName(playlistID) + "\n" + api.PlayListFollowersCount(playlistID) + " people like this playlist!"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Oopsy daisy! there's an issue with lazyfy, pardon!", err)
		os.Exit(1)
	}
}
