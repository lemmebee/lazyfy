package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ehabshaaban/lazyfy/api"
)

var (
	items    []list.Item
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type TrackModel struct {
	list list.Model
}

func (m TrackModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m TrackModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m TrackModel) View() string {
	return docStyle.Render(m.list.View())
}

func NewTracksModel(playlist api.Playlist) TrackModel {
	tracks := api.GetPlaylistTracks(playlist)

	for _, track := range tracks {
		items = append(items, item{
			title: track.Name,
			desc:  api.ConvertTrackArtistListToSingleString(track.Artists[track.Name]),
		})
	}

	return TrackModel{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}
