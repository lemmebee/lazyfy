package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
	"github.com/zmb3/spotify/v2"
)

var (
	playlistItems []list.Item
)

type playlistItem struct {
	id, name string
}

func (p playlistItem) Title() string { return p.name }

// TODO: Description: p.id + # of followers + # of likes + # songs
func (p playlistItem) Description() string { return p.id }
func (p playlistItem) FilterValue() string { return p.name }

type PlaylistModel struct {
	list   list.Model
	choice playlistItem
}

func (m PlaylistModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m PlaylistModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			i, ok := m.list.SelectedItem().(playlistItem)
			if ok {
				m.choice = i

				playlist := api.Playlist{
					ID:   spotify.ID(m.choice.id),
					Name: m.choice.name,
				}
				tracks := NewTracksModel(playlist)
				return tracks, tracks.Init()
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m PlaylistModel) View() string {
	return docStyle.Render(m.list.View())
}

func NewListModel() PlaylistModel {
	playlists := api.GetPlaylists()

	for _, playlist := range playlists {
		playlistItems = append(playlistItems, playlistItem{
			id:   string(playlist.ID),
			name: playlist.Name,
		})
	}

	l := list.New(playlistItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = boldBlueForeground("Look At All Those Playlists!\nwww.youtube.com/watch?v=NsLKQTh-Bqo")
	l.Styles.Title = titleStyle

	return PlaylistModel{
		list: l,
	}
}
