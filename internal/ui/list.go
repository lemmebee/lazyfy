package ui

import (
	"fmt"
	"math/rand"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
	"github.com/muesli/termenv"
)

type ListModel struct {
	playlists []api.Playlist
	cursor    int
	selected  map[int]struct{}
}

func NewListModel(playlists []api.Playlist) ListModel {
	return ListModel{
		playlists: playlists,
		selected:  map[int]struct{}{},
	}
}

func (m ListModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.playlists)-1 {
				m.cursor++
			}
		case "a":
			for i := range m.playlists {
				m.selected[i] = struct{}{}
			}
		case "n":
			for i := range m.selected {
				delete(m.selected, i)
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m ListModel) View() string {
	emojis := []rune("🍦🧋🍡🤠👾😭🦊🐯🦆🥨🎏🍔🍒🍥🎮📦🦁🐶🐸🍕🥐🧲🚒🥇🏆🌽")
	emoji := string(emojis[rand.Intn(len(emojis))])

	s := boldSecondaryForeground("we will make you a playlist on your spotify account with tracks you select " + emoji + "\n\n")

	for i, playlist := range m.playlists {
		s += strconv.Itoa(i)
		line := playlist.Name
		if _, ok := m.selected[i]; ok {
			line = iconSelected + " " + line
		} else {
			line = faint(iconNotSelected + " " + line)
		}
		line += "\n"

		if m.cursor == i {
			nl := ""
			if i > 0 {
				nl = "\n"
			}
			line = nl + boldPrimaryForeground(line) + viewTracks(playlist)
		}

		s += line
	}
	return s
}

func viewTracks(playlist api.Playlist) string {
	var details []string

	var p api.Playlist
	if playlist != p {
		for _, t := range api.GetPlaylistTracks(playlist) {
			details = append(details, fmt.Sprintf("%v", t.Name))
		}
	}

	if len(details) == 0 {
		return ""
	}

	var s string
	for _, d := range details {
		s += "    * " + d + "\n"
	}
	s += "\n"
	return termenv.String(s).Faint().Italic().String()
}
