package ui

import (
	"fmt"
	"math/rand"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
	"github.com/muesli/termenv"
)

// ListModel is the UI in which the user can select which forks should be
// deleted if any, and see details on each of them.
type ListModel struct {
	// mode      string
	playlists []api.Playlist
	cursor    int
	selected  map[int]struct{}
}

// NewListModel creates a new ListModel with the required fields.
func NewListModel(playlists []api.Playlist) ListModel {
	return ListModel{
		playlists: playlists,
		selected:  map[int]struct{}{},
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
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
			// case "d":
			// 	var deleteable []api.Playlist
			// 	for k := range m.selected {
			// 		deleteable = append(deleteable, m.playlists[k])
			// 	}
			// 	dm := NewDeletingModel(m.client, deleteable, m)
			// 	return dm, dm.Init()
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

	// return s + helpView([]helpOption{
	// 	{"up/down", "navigate", false},
	// 	{"space", "toggle selection", false},
	// 	{"d", "delete selected", true},
	// 	{"a", "select all", false},
	// 	{"n", "deselect all", false},
	// 	{"q/esc", "quit", false},
	// })
	return s
}

func viewTracks(playlist api.Playlist) string {
	var details []string

	// empty playlist
	var p api.Playlist

	if playlist != p {
		details = append(details, fmt.Sprintf("%v", api.GetPlaylistTracks(playlist)))
	}
	// if repo.ParentDeleted {
	// 	details = append(details, "Parent repository was deleted")
	// }
	// if repo.ParentDMCATakeDown {
	// 	details = append(details, "Parent repository was taken down by DMCA")
	// }
	// if repo.Private {
	// 	details = append(details, "Is private")
	// }
	// if repo.CommitsAhead > 0 {
	// 	details = append(details, fmt.Sprintf("Has %d commit%s ahead of upstream", repo.CommitsAhead, maybePlural(repo.CommitsAhead)))
	// }
	// if repo.Forks > 0 {
	// 	details = append(details, fmt.Sprintf("Has %d fork%s", repo.Forks, maybePlural(repo.Forks)))
	// }
	// if repo.Stars > 0 {
	// 	details = append(details, fmt.Sprintf("Has %d star%s", repo.Stars, maybePlural(repo.Stars)))
	// }
	// if repo.OpenPRs > 0 {
	// 	details = append(details, fmt.Sprintf("Has %d open PR%s to upstream", repo.OpenPRs, maybePlural(repo.OpenPRs)))
	// }
	// if time.Now().Add(-30 * 24 * time.Hour).Before(repo.LastUpdate) {
	// 	details = append(details, fmt.Sprintf("Was updated recently (%s)", repo.LastUpdate))
	// }

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

func maybePlural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
