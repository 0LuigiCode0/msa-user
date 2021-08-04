package core

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	corehelper "x-msa-core/helper"
	"x-msa-user/core/database"
	"x-msa-user/helper"
	"x-msa-user/hub"

	"github.com/0LuigiCode0/logger"
)

type Server interface {
	Start()
	Close()
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
		s.Close()
		err = fmt.Errorf("db not initialized: %v", err)
		return
	}
	s.hub, err = hub.InitHub(s.db, conf)
	if err != nil {
		s.Close()
		err = fmt.Errorf("hub not initialized: %v", err)
		return
	}
	s.srv.Handler = s.hub.GetHandler()
	s.srv.Addr = fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	logger.Log.Service("server initialized")
	return
}

func (s *server) Start() {
	signal.Notify(corehelper.C, os.Interrupt)

	corehelper.Wg.Add(1)
	go s.loop()

	corehelper.Wg.Add(1)
	go func() {
		defer corehelper.Wg.Done()
		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Log.Service("server stoped")
				corehelper.C <- os.Interrupt
				return
			}
			logger.Log.Errorf("server error: %v", err)
			corehelper.C <- os.Interrupt
			return
		}
	}()

	logger.Log.Service("server started at address:", s.srv.Addr)
	<-corehelper.C
	s.Close()
}

func (s *server) loop() {
	defer corehelper.Wg.Done()
	for {
		select {
		case <-corehelper.Ctx.Done():
			return
		default:
			if err := recover(); err != nil {
				logger.Log.Errorf("critical damage: %v", err)
			}
		}
	}
}

func (s *server) Close() {
	s.srv.Shutdown(corehelper.Ctx)
	if s.hub != nil {
		s.hub.Close()
	}
	if s.db != nil {
		s.db.Close()
	}
	corehelper.CloseCtx()
	corehelper.Wg.Wait()
}
