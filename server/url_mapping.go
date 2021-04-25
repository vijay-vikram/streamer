package server

import "github.com/streamer/controllers/ping"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
