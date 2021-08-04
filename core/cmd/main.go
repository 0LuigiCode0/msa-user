package main

import (
	corehelper "x-msa-core/helper"
	"x-msa-user/core"
	"x-msa-user/helper"

	"github.com/0LuigiCode0/logger"
)

func main() {
	conf := &helper.Config{}
	if err := corehelper.ParseConfig(helper.ConfigDir+helper.ConfigFile, conf); err != nil {
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
