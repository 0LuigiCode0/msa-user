package core

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/0LuigiCode0/msa-user/core/database"
	"github.com/0LuigiCode0/msa-user/helper"
	"github.com/0LuigiCode0/msa-user/hub"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/logger"
)

type Server interface {
	Start()
}

type server struct {
	srv http.Server
	hub hub.Hub
	db  database.DB
}

func InitServer(conf *helper.Config) (S Server, err error) {
	s := &server{}
	S = s
	s.db, err = database.InitDB(conf)
	if err != nil {
		s.close()
		err = fmt.Errorf("db not initialized: %v", err)
		return
	}
	s.hub, err = hub.InitHub(s.db, conf)
	if err != nil {
		s.close()
		err = fmt.Errorf("hub not initialized: %v", err)
		return
	}
	s.srv.Handler = s.hub.GetHandler()
	s.srv.Addr = fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	logger.Log.Service("server initialized")
	return
}

func (s *server) Start() {
	signal.Notify(coreHelper.C, os.Interrupt)

	coreHelper.Wg.Add(1)
	go s.loop()

	coreHelper.Wg.Add(1)
	go func() {
		defer coreHelper.Wg.Done()
		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Log.Service("server stoped")
				coreHelper.C <- os.Interrupt
				return
			}
			logger.Log.Errorf("server error: %v", err)
			coreHelper.C <- os.Interrupt
			return
		}
	}()

	logger.Log.Service("server started at address:", s.srv.Addr)
	<-coreHelper.C
	s.close()
}

func (s *server) loop() {
	defer coreHelper.Wg.Done()
	for {
		select {
		case <-coreHelper.Ctx.Done():
			return
		default:
			if err := recover(); err != nil {
				logger.Log.Errorf("critical damage: %v", err)
			}
		}
	}
}

func (s *server) close() {
	s.srv.Shutdown(coreHelper.Ctx)
	if s.hub != nil {
		s.hub.Close()
	}
	if s.db != nil {
		s.db.Close()
	}
	coreHelper.CloseCtx()
	coreHelper.Wg.Wait()
}
