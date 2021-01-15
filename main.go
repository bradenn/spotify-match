package main

import (
	"spotify-match/config"
	"spotify-match/server"
)

func main() {
	config.Init()
	server.Init()
}
