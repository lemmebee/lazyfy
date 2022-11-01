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
// [x] fix rendering track list issue
// [x] update readme, tui cmds, ci, test, bubbletea, redirect uri
// [x] building
// [x] testing
// [] tagger, releaser - ci
// [x] create new playlist with selected tracks, baf model
// [x] Add create private playlist support
// [x] add tests
// [x] add ci
// [x] update track desc
// [x] change prev from b to backspace + delete track with delete key
// [x] update playlist desc
// [x] select tracks, with highlight
// [x] unselect tracks, with unhighlight
// [] cobra cmd - country specific - defaulted no country
// [] cache auth token
