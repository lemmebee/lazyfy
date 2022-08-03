package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ehabshaaban/lazyfy/internal/ui"
)

var (
	ctx = context.Background()
)

func main() {
	// log.Default().Println(api.DescribePlaylist("7IA6x0ltOvq68X4Lm7MjMN"))
	// log.Default().Println(api.PlayListFollowersCount("7IA6x0ltOvq68X4Lm7MjMN"))
	// log.Default().Println(api.PlaylistName("7IA6x0ltOvq68X4Lm7MjMN"))
	// playlist, msg := api.GetFeaturedPlaylists()
	// log.Default().Println("CEREBRAL::DEBUG:::playlist", playlist)
	// log.Default().Println("CEREBRAL::DEBUG:::msg", msg)
	// log.Default().Println("::::CEREBRAL::::DEBUG::::", api.GetFeaturedPlaylistID())
	// p1 := api.Playlist{
	// 	ID: "p1_7IA6x0ltOvq68X4Lm7MjMN",
	// }
	// p2 := api.Playlist{
	// 	ID: "p2_7IA6x0ltOvq68X4Lm7MjMN",
	// }
	// p3 := api.Playlist{
	// 	ID: "p3_7IA6x0ltOvq68X4Lm7MjMN",
	// }
	// ids := []api.Playlist{p1, p2, p3}
	// p := api.Playlist{ID: id}

	// log.Default().Println("::::CEREBRAL::::DEBUG::::DescribePlaylist", api.DescribePlaylist(p))
	// log.Default().Println("::::CEREBRAL::::DEBUG::::PlayListFollowersCount", api.PlayListFollowersCount(p))
	// log.Default().Println("::::CEREBRAL::::DEBUG::::PlaylistName", api.PlaylistName(p))
	// _, e := api.GetFeaturedPlaylists()
	// log.Default().Println("::::CEREBRAL::::DEBUG::::GetFeaturedPlaylists", e)
	// _, r := api.GetFeaturedPlaylistsWithCountry("FR")
	// log.Default().Println("::::CEREBRAL::::DEBUG::::GetFeaturedPlaylistsWithCountry", r)
	// log.Default().Println("::::CEREBRAL::::DEBUG::::", api.GetPlaylistIDs(api.GetFeaturedPlaylists()))

	// log.Default().Println("::::CEREBRAL::::DEBUG::::", len(api.GetPlaylistIDs(api.GetPlaylists())))
	// log.Default().Println("::::CEREBRAL::::DEBUG::::", api.GetPlaylistIDs(api.GetPlaylists()))

	// log.Default().Println("::::CEREBRAL::::DEBUG::::", len(api.GetPlaylistNames(api.GetPlaylists())))
	// log.Default().Println("::::CEREBRAL::::DEBUG::::", api.GetPlaylistNames(api.GetPlaylists()))

	// log.Default().Println("::::CEREBRAL::::DEBUG::::", len((api.GetPlaylists())))
	// log.Default().Println("::::CEREBRAL::::DEBUG::::", api.GetPlaylists())

	rand.Seed(time.Now().Unix())

	if err := tea.NewProgram(ui.NewInitialModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
