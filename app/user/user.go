package user

import (
	"context"

	"github.com/GopherChat/gopher-server/app"
	"github.com/GopherChat/gopher-server/app/model"
)

type User struct {
	a app.App
}

func New() *User {
	return &User{}
}

func (u *User) Register(ctx context.Context, req *model.RegisterUserReq) (*model.RegisterUserRes, error) {
	panic("not implemented")
}
