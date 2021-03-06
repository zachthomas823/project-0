package config

import "flag"

// FILE is where to find the data to be parsed
const FILE = "./nba_data.csv"

// Player flag for if the user wants an individual player
var Player bool

// Stat flag for if the user wants Stats for every player
var Stat bool

func init() {
	flag.BoolVar(&Player, "player", false, "get player statistics for name passed")
	flag.BoolVar(&Stat, "stat", false, "get the leaders for an individual statistic")
	flag.Parse()
}
