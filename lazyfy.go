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
// [x] add tests
// [x] add ci
// [] update track
// [] change prev from b to back arrow, cmds
// [] create new playlist with selected tracks
// [x] update playlist desc
// [x] select tracks, with highlight
// [x] unselect tracks, with unhighlight
// [] cobra cmd - country specific - defaulted no country
// [] cache auth token
// [] to be created playlist web page
