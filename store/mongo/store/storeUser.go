package store

import (
	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/store/mongo/model"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Store хранилище
type UserStore interface {
	Save(user *model.UserModel) error
	Update(user *model.UserModel) error
	SelectByID(id primitive.ObjectID) (*model.UserModel, error)
	SelectByLogin(login string) (*model.UserModel, error)
}

//Store хранилище
type userStore struct {
	db *mongo.Database
}

func InitUserStore(db *mongo.Database) UserStore {
	return &userStore{db: db}
}

func (s *userStore) Save(user *model.UserModel) error {
	res, err := s.db.Collection(string(helper.CollUsers)).UpdateOne(coreHelper.Ctx, primitive.M{
		"login": user.Login,
	}, primitive.M{
		"$setOnInsert": user,
	}, options.Update().SetUpsert(true))
	if res != nil {
		if id, ok := res.UpsertedID.(primitive.ObjectID); ok {
			user.ID = id
		}
	}
	return err
}

func (s *userStore) Update(user *model.UserModel) error {
	_, err := s.db.Collection(string(helper.CollUsers)).UpdateOne(coreHelper.Ctx, primitive.M{
		"_id": user.ID,
	}, primitive.M{
		"$set": user,
	}, options.Update().SetUpsert(false))
	return err
}

func (s *userStore) SelectByID(id primitive.ObjectID) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := s.db.Collection(string(helper.CollUsers)).FindOne(coreHelper.Ctx, primitive.M{
		"_id": id,
	}).Decode(user)
	return user, err
}

func (s *userStore) SelectByLogin(login string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := s.db.Collection(string(helper.CollUsers)).FindOne(coreHelper.Ctx, primitive.M{
		"login": login,
	}).Decode(user)
	return user, err
}
