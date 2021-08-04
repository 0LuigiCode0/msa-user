package grpcpublic

import (
	"x-msa-core/service/client"
	"x-msa-core/service/server"
	"x-msa-user/store/mongo/model"

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
