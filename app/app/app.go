package app

import (
	"github.com/GopherChat/gopher-server/app"
	"github.com/GopherChat/gopher-server/app/user"
)

type collection struct {
	*user.User
}

type App struct {
	coll collection
}

func New() *App {
	a := &App{}

	a.coll.User = user.New()

	return a
}

func (a *App) User() app.User {
	return a.coll.User
}
