package app

import (
	"context"

	"github.com/GopherChat/gopher-server/app/model"
)

type App interface {
	User() User
}

type User interface {
	Register(context.Context, *model.RegisterUserReq) (*model.RegisterUserRes, error)
}
