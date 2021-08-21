package grpcPublic

import (
	"github.com/0LuigiCode0/msa-user/store/mongo/model"

	"github.com/0LuigiCode0/msa-core/service/client"
	"github.com/0LuigiCode0/msa-core/service/server"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServices interface {
	User() User
}

type userServices struct {
	server.ServiceServer
}

type User interface {
	Error() error

	SelectByID(id primitive.ObjectID) (*model.UserModel, error)
	SelectByLogin(login string) (*model.UserModel, error)
}

type user struct {
	client.ServiceClient
	err error
}
