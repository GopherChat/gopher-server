package sqlstore

import (
	"context"

	"github.com/GopherChat/gopher-server/model"
)

type sqlUserStore struct {
	sqlStore *SqlStore
}

func newSqlUserStore(ss *SqlStore) *sqlUserStore {
	return &sqlUserStore{ss}
}

func (us *sqlUserStore) Find(ctx context.Context, id int64) (*model.User, error) {
	panic("not implemented")
}
