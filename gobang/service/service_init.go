package service

import (
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"gobang/config"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
	addr := config.Config.Get("mongodb.addr").(string)
	db := config.Config.Get("mongodb.db").(string)
	collection := config.Config.Get("mongodb.collection").(string)
	hooker, err := mgorus.NewHooker(addr, db, collection)
	if err == nil {
		logger.Hooks.Add(hooker)
	}
}
