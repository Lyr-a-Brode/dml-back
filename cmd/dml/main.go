package main

import (
	"github.com/Lyr-a-Brode/dml-back/cfg"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	config, err := cfg.Parse()

	if err != nil {
		log.WithField("event", "cfg.parse").Fatal(err)
	}

	if config.EnableDebug {
		log.SetReportCaller(true)
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Test")
}
