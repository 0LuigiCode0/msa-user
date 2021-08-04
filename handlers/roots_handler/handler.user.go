package roots_handler

import (
	"net/http"
	"x-msa-auth/helper"
	corehelper "x-msa-core/helper"
	"x-msa-user/store/mongo/model"

	goutill "github.com/0LuigiCode0/go-utill"
	"github.com/0LuigiCode0/logger"
)

func (h *handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	if _, err := h.Grps().Auth().AuthGuard(r, helper.RoleUser); err != nil {
		logger.Log.Warningf("%v: %v", corehelper.KeyErorrAccessDenied, err)
		h.respError(w, corehelper.ErorrAccessDeniedToken, corehelper.KeyErorrAccessDenied)
		return
	}

	req := &model.UserModel{}
	if err := goutill.JsonParse(r.Body, req); err != nil {
		logger.Log.Warningf("%v json: %v", corehelper.KeyErrorParse, err)
		h.respError(w, corehelper.ErrorParse, corehelper.KeyErrorParse)
		return
	}

	if err := h.MongoStore().UserStore().Save(req); err != nil {
		logger.Log.Warningf("%v user: %v", corehelper.KeyErrorSave, err)
		h.respError(w, corehelper.ErrorSave, corehelper.KeyErrorSave)
		return
	}

	h.respOk(w, "ok")
}
