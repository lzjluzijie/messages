package main

import (
	"net/http"

	"github.com/lzjluzijie/messages/routers"
	macaron "gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Renderer())
	m.Use(macaron.Static("public"))

	m.Get("/", routers.Home)
	m.Get("/message", routers.GetMessage)
	m.Post("/message", routers.PostMessage)

	http.ListenAndServe(":80", m)
}
