package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/api"
)

var trackModels = make(map[string]*trackModel)
var (
	vh int
	vw int
)

type playlist api.Playlist

func (p *playlist) Title() string { return p.Name }

func (p *playlist) Description() string {
	return p.Likes + " likes"
}
func (p *playlist) FilterValue() string { return p.Name }

type PlaylistModel struct {
	list   list.Model
	choice *playlist
}

func (playlistModel *PlaylistModel) Init() tea.Cmd {
	return nil
}

func (playlistModel *PlaylistModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return playlistModel, tea.Quit
		}
		if msg.String() == "enter" {
			p := playlistModel.list.SelectedItem().(*playlist)

			playlistModel.choice = p

			playlist := api.Playlist{
				ID:   playlistModel.choice.ID,
				Name: playlistModel.choice.Name,
			}

			if _, ok := trackModels[playlist.ID]; !ok {
				trackModels[playlist.ID] = NewTracksModel(playlist, playlistModel)
			}
			tracks := trackModels[playlist.ID]
			return tracks, tracks.Init()

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		vh = msg.Height - v
		vw = msg.Width - h
		playlistModel.list.SetSize(vw, vh)
	}

	var cmd tea.Cmd
	playlistModel.list, cmd = playlistModel.list.Update(msg)
	return playlistModel, cmd
}

func (playlistModel PlaylistModel) View() string {
	return docStyle.Render(playlistModel.list.View())
}

func NewPlaylistModel() *PlaylistModel {
	var playlists []list.Item

	for _, p := range api.GetPlaylists() {
		playlists = append(playlists, &playlist{
			ID:    string(p.ID),
			Name:  p.Name,
			Likes: p.Likes,
		})
	}

	l := list.New(playlists, list.NewDefaultDelegate(), 50, 100)
	l.Title = boldBlueForeground("Look At All Those Playlists!\nwww.youtube.com/watch?v=NsLKQTh-Bqo")
	l.Styles.Title = titleStyle

	return &PlaylistModel{
		list: l,
	}
}
