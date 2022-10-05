package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
)

type track api.Track

var SelectedTracks = make(map[string]string)

func (t *track) Title() string {
	if t.Explicit {
		return greenRedForeground(SelectedTracks[t.ID]) + boldRedForeground("E ") + t.Name
	} else {
		return greenRedForeground(SelectedTracks[t.ID]) + t.Name
	}
}

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
			SelectedTracks[t.ID] = star
		}

		if msg.String() == "delete" {
			t := m.list.SelectedItem().(*track)
			SelectedTracks[t.ID] = ""
			delete(SelectedTracks, t.ID)
		}

		if msg.String() == "backspace" {
			var cmd tea.Cmd
			return m.prev, cmd
		}

		if msg.String() == "b" {
			var cmd tea.Cmd
			baf := NewBafModel()
			return baf, cmd
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
			Explicit: t.Explicit,
		})
	}

	l := list.New(tracks, list.NewDefaultDelegate(), 0, 0)
	s := boldBlueForeground(plus)
	l.Title = s + playlist.Name + "\n When you done selecting songs click 'b'"
	l.Styles.Title = titleStyle

	return &trackModel{
		list: l,
		prev: playlistModel,
	}
}
