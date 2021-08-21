package grpcPublic

import (
	"bytes"
	"fmt"

	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/store/mongo/model"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"
	"github.com/0LuigiCode0/msa-core/service/server"

	goutill "github.com/0LuigiCode0/go-utill"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUserServices(service server.ServiceServer) UserServices {
	return &userServices{ServiceServer: service}
}

func (s *userServices) User() User {
	if c, err := s.Services().GetFirstByGroup(coreHelper.User); err != nil {
		return &user{err: err}
	} else {
		return &user{ServiceClient: c}
	}
}

func (u *user) Error() error { return u.err }

func (u *user) SelectByID(id primitive.ObjectID) (*model.UserModel, error) {
	if u.err != nil {
		return nil, u.err
	}

	resp, err := u.Call(&msaService.RequestCall{
		FuncName: helper.SelectByID,
		Args:     map[string]string{"user_id": id.Hex()},
	})
	if err != nil {
		return nil, fmt.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
	}
	out := &model.UserModel{}
	if err := goutill.JsonParse(bytes.NewReader(resp.Result), out); err != nil {
		return nil, fmt.Errorf("%v json: %v", coreHelper.KeyErrorParse, err)
	}
	return out, nil
}

func (u *user) SelectByLogin(login string) (*model.UserModel, error) {
	if u.err != nil {
		return nil, u.err
	}

	resp, err := u.Call(&msaService.RequestCall{
		FuncName: helper.SelectByLogin,
		Args:     map[string]string{"login": login},
	})
	if err != nil {
		return nil, fmt.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
	}
	out := &model.UserModel{}
	if err := goutill.JsonParse(bytes.NewReader(resp.Result), out); err != nil {
		return nil, fmt.Errorf("%v json: %v", coreHelper.KeyErrorParse, err)
	}
	return out, nil
}
