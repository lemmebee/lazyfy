// package ui

// import (
// 	"fmt"
// 	"math/rand"
// 	"strconv"

// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/ehabshaaban/lazyfy/api"
// 	"github.com/muesli/termenv"
// )

// type ListModel struct {
// 	playlists []api.Playlist
// 	cursor    int
// 	selected  map[int]struct{}
// }

// func NewListModel(playlists []api.Playlist) ListModel {
// 	return ListModel{
// 		playlists: playlists,
// 		selected:  map[int]struct{}{},
// 	}
// }

// func (m ListModel) Init() tea.Cmd {
// 	return tea.EnterAltScreen
// }

// func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q", "esc":
// 			return m, tea.Quit
// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}
// 		case "down", "j":
// 			if m.cursor < len(m.playlists)-1 {
// 				m.cursor++
// 			}
// 		case "a":
// 			for i := range m.playlists {
// 				m.selected[i] = struct{}{}
// 			}
// 		case "enter":
// 			for _, playlist := range m.playlists {
// 				tracks := NewTracksModel(playlist)
// 				return tracks, tracks.Init()
// 			}
// 		case "n":
// 			for i := range m.selected {
// 				delete(m.selected, i)
// 			}
// 		case " ":
// 			_, ok := m.selected[m.cursor]
// 			if ok {
// 				delete(m.selected, m.cursor)
// 			} else {
// 				m.selected[m.cursor] = struct{}{}
// 			}
// 		}
// 	}
// 	return m, nil
// }

// func (m ListModel) View() string {
// 	emojis := []rune("🍦🧋🍡🤠👾😭🦊🐯🦆🥨🎏🍔🍒🍥🎮📦🦁🐶🐸🍕🥐🧲🚒🥇🏆🌽")
// 	emoji := string(emojis[rand.Intn(len(emojis))])

// 	s := boldSecondaryForeground("we will make you a playlist on your spotify account with tracks you select " + emoji + "\n\n")

// 	for i, playlist := range m.playlists {
// 		s += strconv.Itoa(i)
// 		line := playlist.Name
// 		if _, ok := m.selected[i]; ok {
// 			line = iconSelected + " " + line
// 		} else {
// 			line = faint(iconNotSelected + " " + line)
// 		}
// 		line += "\n"

// 		if m.cursor == i {
// 			nl := ""
// 			if i > 0 {
// 				nl = "\n"
// 			}
// 			line = nl + boldPrimaryForeground(line) + viewTracks(playlist)
// 		}

// 		s += line
// 	}
// 	return s
// }

// func viewTracks(playlist api.Playlist) string {
// 	var details []string

// 	var p api.Playlist
// 	if playlist != p {
// 		// for i, t := range api.GetPlaylistTracks(playlist) {
// 		// 	// details = append(details, fmt.Sprintf("%v", t.Name))
// 		// 	details = append(details, fmt.Sprintf("%v::::%v", i, t.Name))
// 		// }
// 		followerCount := api.GetPlayListFollowersCount(playlist)
// 		trackCount := api.GetPlayListTrackCount(playlist)
// 		details = append(details, fmt.Sprintf("%v Followers %v Tracks", followerCount, trackCount))
// 	}

// 	if len(details) == 0 {
// 		return ""
// 	}

// 	var s string
// 	for _, d := range details {
// 		s += "    * " + d + "\n"
// 	}
// 	s += "\n"
// 	return termenv.String(s).Faint().Italic().String()
// }

package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zmb3/spotify/v2"

	"github.com/ehabshaaban/lazyfy/api"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

// type item string
type item spotify.ID

// type item struct {
// 	name string
// 	id   spotify.ID
// }

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	list     list.Model
	items    []item
	choice   item
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i

				playlist := api.Playlist{
					ID: spotify.ID(m.choice),
				}
				tracks := NewTracksModel(playlist)
				return tracks, tracks.Init()
			}
			// // return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// if m.choice != "" {
	// 	return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	// }
	// if m.quitting {
	// 	return quitTextStyle.Render("Not hungry? That’s cool.")
	// }
	return m.list.View()
}

func NewListModel() model {
	playlists := api.GetPlaylists()

	for _, playlist := range playlists {
		items = append(items, item(playlist.ID))
	}

	// items := []list.Item{
	// 	item("Ramen"),
	// 	item("Tomato Soup"),
	// 	item("Hamburgers"),
	// 	item("Cheeseburgers"),
	// 	item("Currywurst"),
	// 	item("Okonomiyaki"),
	// 	item("Pasta"),
	// 	item("Fillet Mignon"),
	// 	item("Caviar"),
	// 	item("Just Wine"),
	// }

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Look At All Those Playlists!\nwww.youtube.com/watch?v=NsLKQTh-Bqo"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		list: l,
	}

	// m := model{list: l}

	// if err := tea.NewProgram(m).Start(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }
}
