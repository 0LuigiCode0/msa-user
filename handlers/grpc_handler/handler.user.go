package grpc_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"x-msa-core/grpc/msa_service"
	corehelper "x-msa-core/helper"
	"x-msa-user/helper"
	"x-msa-user/store/mongo/model"

	"github.com/0LuigiCode0/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) call(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error) {
	var out interface{}
	var err error

	switch req.FuncName {
	case helper.SelectByID:
		out, err = h.selectByID(req)
	case helper.SelectByLogin:
		out, err = h.selectByLogin(req)
	default:
		logger.Log.Warningf("%v func -> %v", corehelper.KeyErrorNotFound, req.FuncName)
		return nil, fmt.Errorf("%v func -> %v", corehelper.KeyErrorNotFound, req.FuncName)
	}
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(out)
	if err != nil {
		logger.Log.Warningf("%v json: %v", corehelper.KeyErrorParse, err)
		return nil, err
	}
	return &msa_service.ResponseCall{Result: data}, nil
}

func (h *handler) selectByID(req *msa_service.RequestCall) (*model.UserModel, error) {
	if v, ok := req.Args["user_id"]; ok {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			logger.Log.Warningf("%v id: %v", corehelper.KeyErrorParse, err)
			return nil, fmt.Errorf("%v id: %v", corehelper.KeyErrorParse, err)
		}
		resp, err := h.MongoStore().UserStore().SelectByID(id)
		if err != nil {
			logger.Log.Errorf("%v user: %v", corehelper.KeyErrorNotFound, err)
			return nil, err
		}
		return resp, nil
	}
	logger.Log.Warningf("%v id", corehelper.ErrorNotFound)
	return nil, fmt.Errorf("%v id", corehelper.ErrorNotFound)
}

func (h *handler) selectByLogin(req *msa_service.RequestCall) (*model.UserModel, error) {
	if v, ok := req.Args["login"]; ok {
		resp, err := h.MongoStore().UserStore().SelectByLogin(v)
		if err != nil {
			logger.Log.Errorf("%v user: %v", corehelper.KeyErrorNotFound, err)
			return nil, err
		}
		return resp, nil
	}
	logger.Log.Warningf("%v login", corehelper.ErrorNotFound)
	return nil, fmt.Errorf("%v login", corehelper.ErrorNotFound)
}
