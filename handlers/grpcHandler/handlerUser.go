package grpcHandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/store/mongo/model"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) call(ctx context.Context, req *msaService.RequestCall) (*msaService.ResponseCall, error) {
	var out interface{}
	var err error

	switch req.FuncName {
	case helper.SelectByID:
		out, err = h.selectByID(req)
	case helper.SelectByLogin:
		out, err = h.selectByLogin(req)
	default:
		logger.Log.Warningf("%v func -> %v", coreHelper.KeyErrorNotFound, req.FuncName)
		return nil, fmt.Errorf("%v func -> %v", coreHelper.KeyErrorNotFound, req.FuncName)
	}
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(out)
	if err != nil {
		logger.Log.Warningf("%v json: %v", coreHelper.KeyErrorParse, err)
		return nil, err
	}
	return &msaService.ResponseCall{Result: data}, nil
}

func (h *handler) selectByID(req *msaService.RequestCall) (*model.UserModel, error) {
	if v, ok := req.Args["user_id"]; ok {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			logger.Log.Warningf("%v id: %v", coreHelper.KeyErrorParse, err)
			return nil, fmt.Errorf("%v id: %v", coreHelper.KeyErrorParse, err)
		}
		resp, err := h.MongoStore().UserStore().SelectByID(id)
		if err != nil {
			logger.Log.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
			return nil, err
		}
		return resp, nil
	}
	logger.Log.Warningf("%v id", coreHelper.ErrorNotFound)
	return nil, fmt.Errorf("%v id", coreHelper.ErrorNotFound)
}

func (h *handler) selectByLogin(req *msaService.RequestCall) (*model.UserModel, error) {
	if v, ok := req.Args["login"]; ok {
		resp, err := h.MongoStore().UserStore().SelectByLogin(v)
		if err != nil {
			logger.Log.Errorf("%v user: %v", coreHelper.KeyErrorNotFound, err)
			return nil, err
		}
		return resp, nil
	}
	logger.Log.Warningf("%v login", coreHelper.ErrorNotFound)
	return nil, fmt.Errorf("%v login", coreHelper.ErrorNotFound)
}
