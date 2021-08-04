package roots_handler

import (
	"encoding/json"
	"net/http"
	corehelper "x-msa-core/helper"
	"x-msa-user/handlers/roots_handler/roots_helper"
	"x-msa-user/helper"
	"x-msa-user/hub/hub_helper"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hub_helper.HelperForHandler
}

func InitHandler(hub hub_helper.HelperForHandler, conf *helper.HandlerConfig) (H roots_helper.Handler, err error) {
	h := &handler{HelperForHandler: hub}
	H = h

	hUser := h.Router().PathPrefix("/user").Subrouter()
	hUser.HandleFunc("/create", h.UserCreate)

	h.SetHandler(applyCORS(h.Router()))
	return
}

func (h *handler) respOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &corehelper.ResponseModel{
		Success: true,
		Result:  data,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		logger.Log.Warningf(corehelper.KeyErrorParse+": josn: %v", err)
		h.respError(w, corehelper.ErrorParse, corehelper.KeyErrorParse+": josn")
		return
	}
	_, err = w.Write(buf)
	if err != nil {
		logger.Log.Warningf(corehelper.KeyErrorWrite+": response: %v", err)
		h.respError(w, corehelper.ErrorWrite, corehelper.KeyErrorWrite+": response")
		return
	}
}

func (h *handler) respError(w http.ResponseWriter, code corehelper.ErrCode, msg string) {
	w.Header().Set("Content-Type", "application/json")
	resp := &corehelper.ResponseModel{
		Success: false,
		Result: &corehelper.ResponseError{
			Code: code,
			Msg:  msg,
		},
	}
	buf, _ := json.Marshal(resp)
	_, err := w.Write(buf)
	if err != nil {
		logger.Log.Warningf(corehelper.KeyErrorWrite+": response: %v", err)
	}
}
