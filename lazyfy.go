package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/internal/ui"
)

var (
	ctx = context.Background()
)

func main() {
	rand.Seed(time.Now().Unix())

	if err := tea.NewProgram(ui.NewInitialModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
