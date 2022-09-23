package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
)

type trackItem struct {
	name, artists string
}

func (t trackItem) Title() string { return t.name }

// TODO: Description: i.artists + track ablum + track duration + explicit
func (t trackItem) Description() string { return t.artists }
func (t trackItem) FilterValue() string { return t.name }

type TrackModel struct {
	list list.Model
	prev PlaylistModel
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

func (m TrackModel) View() string {
	return docStyle.Render(m.list.View())
}

func NewTracksModel(playlist api.Playlist, playlistModel PlaylistModel) TrackModel {
	var trackItems []list.Item

	tracks := api.GetPlaylistTracks(&playlist)

	for _, track := range tracks {
		trackItems = append(trackItems, trackItem{
			name:    track.Name,
			artists: api.ConvertTrackArtistListToSingleString(track.Artists[track.Name]),
		})
	}

	l := list.New(trackItems, list.NewDefaultDelegate(), 0, 0)
	s := boldRedForeground(star)
	l.Title = s + playlist.Name
	l.Styles.Title = titleStyle

	return TrackModel{
		list: l,
		prev: playlistModel,
	}
}
