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
	if err := tea.NewProgram(ui.NewListModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running lazyfy!:", err)
		os.Exit(1)
	}
}
