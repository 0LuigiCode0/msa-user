package hub

import (
	"fmt"
	"net/http"
	corehelper "x-msa-core/helper"
	"x-msa-user/core/database"
	"x-msa-user/handlers/grpc_handler"
	"x-msa-user/handlers/grpc_handler/grpc_helper"
	"x-msa-user/handlers/roots_handler"
	"x-msa-user/handlers/roots_handler/roots_helper"
	"x-msa-user/helper"
	"x-msa-user/hub/hub_helper"
	"x-msa-user/store/mongo/model"

	"github.com/0LuigiCode0/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_roots = "roots"
	_wss   = "wss"
	_grpc  = "grpc"
)

type Hub interface {
	GetHandler() http.Handler
	Close()
}

type hub struct {
	helper hub_helper.Helper
	database.DB
	router  *mux.Router
	handler http.Handler
	config  *helper.Config

	_roots roots_helper.Handler
	_grpc  grpc_helper.Handler
}

func InitHub(db database.DB, conf *helper.Config) (H Hub, err error) {
	hh := &hub{
		DB:     db,
		router: mux.NewRouter(),
		config: conf,
	}
	H = hh
	hh.SetHandler(hh.router)

	hh.helper = hub_helper.InitHelper(hh)

	if err = hh.intiDefault(); err != nil {
		logger.Log.Warningf("initializing default is failed: %v", err)
		err = fmt.Errorf("handler not initializing: %v", err)
		return
	}
	logger.Log.Service("initializing default")

	if v, ok := conf.Handlers[_roots]; ok {
		hh._roots, err = roots_handler.InitHandler(hh, v)
		if err != nil {
			err = fmt.Errorf("handler %q not initializing: %v", _roots, err)
			return
		}
		logger.Log.Servicef("handler %q initializing", _roots)
	} else {
		err = fmt.Errorf("config %q not found", _roots)
		return
	}

	if v, ok := conf.Handlers[_grpc]; ok {
		hh._grpc, err = grpc_handler.InitHandler(hh, v)
		if err != nil {
			err = fmt.Errorf("handler %q not initializing: %v", _grpc, err)
			return
		}
		logger.Log.Servicef("handler %q initializing", _grpc)
	} else {
		err = fmt.Errorf("config %q not found", _grpc)
		return
	}

	logger.Log.Service("handler initializing")
	return
}

func (h *hub) Config() *helper.Config      { return h.config }
func (h *hub) Helper() hub_helper.Helper   { return h.helper }
func (h *hub) Router() *mux.Router         { return h.router }
func (h *hub) GetHandler() http.Handler    { return h.handler }
func (h *hub) SetHandler(hh http.Handler)  { h.handler = hh }
func (h *hub) Roots() roots_helper.Handler { return h._roots }
func (h *hub) Grps() grpc_helper.MSA       { return h._grpc }
func (h *hub) Close() {
	if h._grpc != nil {
		h._grpc.Close()
	}
}

func (h *hub) intiDefault() error {
	// Users collection create
	if _, err := h.Mongo().Collection(string(helper.CollUsers)).Indexes().CreateMany(
		corehelper.Ctx,
		[]mongo.IndexModel{
			{
				Keys:    primitive.M{"login": 1},
				Options: options.Index().SetUnique(true),
			},
		},
	); err != nil {
		logger.Log.Errorf("created collection %v is failed : %v", helper.CollUsers, err)
	}
	logger.Log.Servicef("collection %v is created", helper.CollUsers)

	for _, user := range h.Config().Admins {
		pwd, _ := corehelper.Hash(user.Password)
		newUser := &model.UserModel{Login: user.Login, Password: pwd}
		h.MongoStore().UserStore().Save(newUser)
		if newUser.ID.IsZero() {
			h.MongoStore().UserStore().Update(newUser)
		}
	}

	return nil
}
