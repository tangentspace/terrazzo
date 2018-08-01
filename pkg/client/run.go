package client

import (
	log "github.com/sirupsen/logrus"
	"github.com/tangentspace/terrazzo/pkg/config"
	"github.com/tangentspace/terrazzo/pkg/controller"
	"github.com/tangentspace/terrazzo/pkg/handlers"
)

func Run(conf *config.Config) {
	eventHandler := new(handlers.Default)
	if err := eventHandler.Init(conf); err != nil {
		log.Fatal(err)
	}
	controller.Start(conf, eventHandler)
}
