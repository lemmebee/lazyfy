package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
)

type track api.Track

var selectedTracks = make(map[string]string)

func (t *track) Title() string {
	return boldRedForeground(selectedTracks[t.ID]) + t.Name
}

// TODO: Description: i.artists + track ablum + track duration + isExplicit
func (t *track) Description() string {
	return t.Duration + ", " + api.ConvertTrackArtistListToSingleString(t.Artists[t.Name])
}
func (t *track) FilterValue() string { return t.Name }

type trackModel struct {
	list list.Model
	prev *PlaylistModel
}

func (m *trackModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *trackModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == " " {
			t := m.list.SelectedItem().(*track)
			selectedTracks[t.ID] = star
		}

		if msg.String() == "r" {
			t := m.list.SelectedItem().(*track)
			selectedTracks[t.ID] = ""
			delete(selectedTracks, t.ID)
		}

		if msg.String() == "b" {
			var cmd tea.Cmd
			return m.prev, cmd
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *trackModel) View() string {
	return docStyle.Render(m.list.View())
}

func NewTracksModel(playlist api.Playlist, playlistModel *PlaylistModel) *trackModel {
	var tracks []list.Item

	for _, t := range api.GetPlaylistTracks(&playlist) {
		tracks = append(tracks, &track{
			ID:       t.ID,
			Name:     t.Name,
			Artists:  t.Artists,
			Duration: t.Duration,
		})
	}

	l := list.New(tracks, list.NewDefaultDelegate(), 0, 0)
	s := boldBlueForeground(plus)
	l.Title = s + playlist.Name
	l.Styles.Title = titleStyle

	return &trackModel{
		list: l,
		prev: playlistModel,
	}
}
