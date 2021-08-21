package model

import (
	"github.com/0LuigiCode0/msa-auth/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserModel модель пользователя в БД
type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Login    string             `bson:"login" json:"login"`
	Password string             `bson:"password" json:"password"`
	Role     helper.Role        `bson:"role" json:"role"`
}
