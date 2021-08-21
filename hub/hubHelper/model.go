package hubHelper

import (
	"net/http"

	"github.com/0LuigiCode0/msa-user/core/database"
	"github.com/0LuigiCode0/msa-user/handlers/grpcHandler/grpcHelper"
	"github.com/0LuigiCode0/msa-user/handlers/rootsHandler/rootsHelper"
	"github.com/0LuigiCode0/msa-user/helper"

	"github.com/gorilla/mux"
)

type Helper interface {
}

type HelperForHandler interface {
	database.DBForHandler
	Helper() Helper
	Config() *helper.Config
	Router() *mux.Router
	SetHandler(hh http.Handler)
	Grps() grpcHelper.MSA
}

type HandlerForHelper interface {
	database.DBForHandler
	Roots() rootsHelper.Handler
	Grps() grpcHelper.MSA
	Config() *helper.Config
}

type help struct {
	HandlerForHelper
}
