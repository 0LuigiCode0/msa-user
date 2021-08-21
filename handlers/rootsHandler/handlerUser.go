package rootsHandler

import (
	"net/http"

	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/store/mongo/model"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
	authHelper "github.com/0LuigiCode0/msa-auth/helper"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"
)

func (h *handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	if _, err := h.Grps().Auth().AuthGuard(r, authHelper.RoleUser); err != nil {
		logger.Log.Warningf("%v: %v", coreHelper.KeyErorrAccessDenied, err)
		h.respError(w, coreHelper.ErorrAccessDeniedToken, coreHelper.KeyErorrAccessDenied)
		return
	}

	req := &model.UserModel{}
	if err := goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf("%v json: %v", coreHelper.KeyErrorParse, err)
		h.respError(w, coreHelper.ErrorParse, coreHelper.KeyErrorParse)
		return
	}

	req.Password = coreHelper.Hash(req.Password, helper.Secret)

	if err := h.MongoStore().UserStore().Save(req); err != nil {
		logger.Log.Warningf("%v user: %v", coreHelper.KeyErrorSave, err)
		h.respError(w, coreHelper.ErrorSave, coreHelper.KeyErrorSave)
		return
	}

	h.respOk(w, "ok")
}
