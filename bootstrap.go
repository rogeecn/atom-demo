package main

import (
	"atom/modules"

	"github.com/rogeecn/atom/providers"
	"github.com/rogeecn/atom/providers/config"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/http/gin"
	"github.com/rogeecn/atom/providers/log"
)

func bootstrap() {
	config := config.AutoLoad()

	p := providers.New(providers.WithConfig(config), providers.WithConfigPrefix(log.DefaultPrefix))
	if err := log.Provide(p); err != nil {
		log.Fatal("provide http service failed, err: ", err)
	}

	p = providers.New(providers.WithConfig(config), providers.WithConfigPrefix(http.DefaultPrefix))
	if err := gin.Provide(p); err != nil {
		log.Fatal("provide http service failed, err: ", err)
	}

	if err := modules.Provide(); err != nil {
		log.Fatal("provide modules failed, err: ", err)
	}
}
