package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
	"github.com/zmb3/spotify/v2"
)

var (
	playlistItems []list.Item
	state         bool = false
	foo           TrackModel
	trackModels   map[spotify.ID]TrackModel = make(map[spotify.ID]TrackModel)
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

func (playlistModel PlaylistModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (playlistModel PlaylistModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return playlistModel, tea.Quit
		}
		if msg.String() == "enter" {
			i := playlistModel.list.SelectedItem().(playlistItem)

			playlistModel.choice = i

			playlist := api.Playlist{
				ID:   spotify.ID(playlistModel.choice.id),
				Name: playlistModel.choice.name,
			}

			if _, ok := trackModels[playlist.ID]; !ok {
				trackModels[playlist.ID] = NewTracksModel(playlist, playlistModel)
			}
			tracks := trackModels[playlist.ID]
			return tracks, tracks.Init()

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		playlistModel.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	playlistModel.list, cmd = playlistModel.list.Update(msg)
	return playlistModel, cmd
}

func (playlistModel PlaylistModel) View() string {
	return docStyle.Render(playlistModel.list.View())
}

func NewPlaylistModel() PlaylistModel {
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
