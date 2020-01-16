package main

import (
	"fmt"

	"github.com/project-0/config"
	"github.com/project-0/stats"
)

func main() {
	if config.Player {
		stats.PlayerStats()
	} else if config.Stat {
		stats.StatLeader()
	} else {
		fmt.Println("Please use tag -player \"Player Name\" \nor tag -stat \"Stat abrev\"")
	}
}
