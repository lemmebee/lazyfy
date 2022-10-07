package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/log"
)

const (
	normalMode = iota
	isPlaylistPublicMode
	byeByeMode
)

var (
	playlistName     string
	isPlaylistPublic bool
	items            = []list.Item{
		item("Yes"),
		item("No"),
	}
	l            = list.New(items, itemDelegate{}, defaultWidth, listHeight)
	listHeight   = 14
	defaultWidth = 20
)

type item string

type itemDelegate struct{}

func (i item) FilterValue() string { return "" }

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
	textInput textinput.Model
	mode      int
	list      list.Model
	items     []item
	choice    string
}

func NewBafModel() model {
	ti := textinput.New()
	ti.Placeholder = "lazyfy is the shit! Thanks for making me listen to only good music playlist?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		mode:      normalMode,
		list:      l,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case normalMode:
		return m.normalUpdate(msg)
	case isPlaylistPublicMode:
		return m.isPlaylistPublicUpdate(msg)
	}

	return m, nil
}

func (m model) View() string {
	switch m.mode {
	case normalMode:
		return m.normalView()
	case isPlaylistPublicMode:
		return m.isPlaylistPublicView()

	}

	return ""
}

func (m model) normalUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			playlistName = m.textInput.Value()
			log.Log(m.mode)
			log.Log(playlistName)
			m.mode = isPlaylistPublicMode
		}
	}

	m.textInput, _ = m.textInput.Update(msg)
	return m, nil
}

func (m model) isPlaylistPublicUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				if m.choice == "Yes" {
					isPlaylistPublic = true
				} else {
					isPlaylistPublic = false
				}
			}
			log.Log(m.mode)
			log.Log(isPlaylistPublic)
			m.mode = byeByeMode
			log.Log(m.mode)
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) normalView() string {
	return fmt.Sprintf(
		"lazyfy is creating a playlist with the songs you selected\nWhat’s your playlist name gonna be like?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func (m model) isPlaylistPublicView() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s", m.choice))
	}
	return "\n" + m.list.View()
}