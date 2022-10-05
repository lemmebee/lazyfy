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
// [] create new playlist with selected tracks, baf model
// [] Add private playlist support in LazyfyForUser()
// [] to be created playlist web page
// [] update readme, tui cmds, ci, test, bubbletea, redirect uri
// [x] add tests
// [x] add ci
// [x] update track desc
// [x] change prev from b to backspace + delete track with delete key
// [x] update playlist desc
// [x] select tracks, with highlight
// [x] unselect tracks, with unhighlight
// [] cobra cmd - country specific - defaulted no country
// [] cache auth token
