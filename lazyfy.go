package main

import (
	"log"

	"github.com/ehabshaaban/lazyfy/api"
)

func main() {
	log.Default().Println(api.DescribePlaylist("7IA6x0ltOvq68X4Lm7MjMN"))
	log.Default().Println(api.PlayListFollowersCount("7IA6x0ltOvq68X4Lm7MjMN"))
	log.Default().Println(api.PlaylistName("7IA6x0ltOvq68X4Lm7MjMN"))
}
