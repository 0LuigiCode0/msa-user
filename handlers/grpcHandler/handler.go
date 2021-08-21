package grpcHandler

import (
	"fmt"
	"os"

	"github.com/0LuigiCode0/msa-auth/handlers/grpcHandler/grpcPublic"

	"github.com/0LuigiCode0/msa-user/handlers/grpcHandler/grpcHelper"
	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/hub/hubHelper"

	"github.com/0LuigiCode0/msa-core/grpc/msaObserver"
	coreHelper "github.com/0LuigiCode0/msa-core/helper"
	"github.com/0LuigiCode0/msa-core/service/server"

	"github.com/0LuigiCode0/logger"
)

type handler struct {
	hubHelper.HelperForHandler
	server.ServiceServer

	grpcPublic.AuthServices
}

func InitHandler(hub hubHelper.HelperForHandler, conf *helper.HandlerConfig) (H grpcHelper.Handler, err error) {
	h := &handler{
		HelperForHandler: hub,
		ServiceServer:    server.NewServiceServer(conf.Key, fmt.Sprintf("%v:%v", conf.Host, conf.Port)),
	}
	H = h

	h.AuthServices = grpcPublic.NewUserServices(h)

	h.SetCall(h.call)

	coreHelper.Wg.Add(1)
	go h.start()

	if err = h.initDependents(); err != nil {
		logger.Log.Errorf("init dependents error: %v", err)
		return
	}

	logger.Log.Servicef("gserver started at address: %v", fmt.Sprintf("%v:%v", conf.Host, conf.Port))
	return
}

func (h *handler) start() {
	defer coreHelper.Wg.Done()

	if err := h.Start(); err != nil {
		logger.Log.Errorf("canot start gserver %v: %v", h.GetAddr(), err)
		coreHelper.C <- os.Interrupt
		return
	}
}

func (h *handler) initDependents() error {
	for _, v := range h.Config().Observers {
		coreHelper.Wg.Add(1)
		go h.addObserver(v.Key, fmt.Sprintf("%v:%v", v.Host, v.Port))
	}
	for _, v := range h.Config().Services {
		coreHelper.Wg.Add(1)
		go h.AddService(v.Key, fmt.Sprintf("%v:%v", v.Host, v.Port), v.Group)
	}
	return nil
}

func (h *handler) addObserver(key, addr string) {
	defer coreHelper.Wg.Done()

	if err := h.Observers().Add(key, addr); err != nil {
		logger.Log.Warningf("canot added observer key %v: %v", key, err)
		return
	}
	observer, err := h.Observers().Get(key)
	if err != nil {
		logger.Log.Warningf("canot find observer key %v: %v", key, err)
		return
	}
	res, err := observer.PushFirst(&msaObserver.RequestPushFirst{})
	if err != nil {
		logger.Log.Warningf("canot request in observer %v: %v", key, err)
		return
	}
	fmt.Println(res)
}

func (h *handler) AddService(key, addr string, group coreHelper.GroupsType) {
	defer coreHelper.Wg.Done()

	if err := h.Services().Add(key, addr, group); err != nil {
		logger.Log.Warningf("canot added service key %v addr %v: %v", key, addr, err)
		return
	}
	logger.Log.Servicef("added service key %v, addr%v", key, addr)
}

func (h *handler) DeleteService(key string, group coreHelper.GroupsType) error {
	if err := h.Services().Delete(group, key); err != nil {
		return err
	}
	return nil
}
