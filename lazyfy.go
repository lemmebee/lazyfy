package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/internal/ui"
)

var (
	ctx = context.Background()
)

func main() {
	if err := tea.NewProgram(ui.NewInitialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running lazyfy!:", err)
		os.Exit(1)
	}

	// playlist := api.Playlist{
	// 	ID: "3CddHe2KUSiOtXxFdwrUkL",
	// }

	// tracks := api.GetPlaylistTracks(playlist)

	// fmt.Println(tracks[1])

	// // track name
	// trackName := tracks[1].Name
	// fmt.Println(trackName)

	// // track artist
	// fmt.Println(tracks[1].Artists[trackName])

	// ids := api.GetPlaylistIDs(api.GetPlaylists())
	// fmt.Println(ids)
	// playlist := api.Playlist{
	// 	ID: spotify.ID(ids[1]),
	// }
	// fmt.Println(playlist)
	// tracks := api.GetPlaylistTracks(playlist)
	// fmt.Println(tracks)

}
