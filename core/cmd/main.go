package main

import (
	"github.com/0LuigiCode0/msa-user/core"
	"github.com/0LuigiCode0/msa-user/helper"

	coreHelper "github.com/0LuigiCode0/msa-core/helper"

	"github.com/0LuigiCode0/logger"
)

func main() {
	conf := &helper.Config{}
	if err := coreHelper.ParseConfig(helper.ConfigDir+helper.ConfigFile, conf); err != nil {
		logger.Log.Errorf("config parse invalid: %v", err)
		return
	}
	srv, err := core.InitServer(conf)
	if err != nil {
		logger.Log.Errorf("server not initialized: %v", err)
		return
	}
	srv.Start()
}
