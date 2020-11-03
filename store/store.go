package store

import (
	"context"

	"github.com/GopherChat/gopher-server/model"
)

type Store interface {
	User() UserStore
}

type UserStore interface {
	Find(context.Context, int64) (*model.User, error)
}
