package ui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ehabshaaban/lazyfy/api"
)

type initialModel struct {
	playlistIDs []string
	index       int
	width       int
	height      int
	spinner     spinner.Model
	progress    progress.Model
	done        bool
}

var (
	currentPlaylistIdStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("239"))
	doneStyle              = lipgloss.NewStyle().Margin(1, 2)
	checkMark              = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func NewInitialModel() initialModel {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	return initialModel{
		playlistIDs: api.GetPlaylistIDs(api.GetPlaylists()),
		spinner:     s,
		progress:    p,
	}
}

func (m initialModel) Init() tea.Cmd {
	return tea.Batch(getPlaylistID(m.playlistIDs[m.index]), getPlaylists(), m.spinner.Tick)
}

func (m initialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case getPlaylistsMsg:
		list := NewListModel(msg.playlists)
		return list, list.Init()
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	case getPlaylistIDMsg:
		if m.index >= len(m.playlistIDs)-1 {
			// Everything's been installed. We're done!
			m.done = true
			// return m, tea.Quit
			return m, getPlaylists()
		}

		// Update progress bar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.playlistIDs)-1))

		m.index++
		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s %s", m.index, checkMark, m.playlistIDs[m.index]), // print success message
			getPlaylistID(m.playlistIDs[m.index]),                              // fetch and print the next playlist ID
		)
	}
	return m, nil
}

func (m initialModel) View() string {
	n := len(m.playlistIDs)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Fetched %d playlists.\n", n))
	}

	playlistCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+playlistCount))

	playlistID := currentPlaylistIdStyle.Render(m.playlistIDs[m.index])
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Fetching " + playlistID)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+playlistCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + playlistCount
}

type getPlaylistIDMsg string

func getPlaylistID(playlistID string) tea.Cmd {
	d := time.Millisecond * time.Duration(rand.Intn(500))
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return getPlaylistIDMsg(playlistID)
	})
}

type getPlaylistsMsg struct {
	playlists []api.Playlist
}

func getPlaylists() tea.Cmd {
	return func() tea.Msg {
		playlists := api.GetPlaylists()
		return getPlaylistsMsg{playlists}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
