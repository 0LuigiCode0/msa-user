package hub_helper

import (
	"net/http"
	"x-msa-user/core/database"
	"x-msa-user/handlers/grpc_handler/grpc_helper"
	"x-msa-user/handlers/roots_handler/roots_helper"
	"x-msa-user/helper"

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
	Grps() grpc_helper.MSA
}

type HandlerForHelper interface {
	database.DBForHandler
	Roots() roots_helper.Handler
	Grps() grpc_helper.MSA
	Config() *helper.Config
}

type help struct {
	HandlerForHelper
}
