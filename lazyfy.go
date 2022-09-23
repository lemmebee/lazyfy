package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/internal/ui"
)

func main() {
	if err := tea.NewProgram(ui.NewPlaylistModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running lazyfy!:", err)
		os.Exit(1)
	}
}

// TODO:
// [] add tests
// [] add ci
// [] select tracks, with highlight
// [] unselect tracks, with unhighlight
// [] cache auth token
// [] to be created playlist web page
// [] create new playlist with selected tracks
