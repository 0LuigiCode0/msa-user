package mongoStore

import (
	"x-msa-user/store/mongo/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	UserStore() store.UserStore
}

type s struct {
	user store.UserStore
}

func InitStore(db *mongo.Database) Store {
	return &s{
		user: store.InitUserStore(db),
	}
}

func (s *s) UserStore() store.UserStore { return s.user }
