package service

import (
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
	hooker, err := mgorus.NewHooker("150.158.104.248:27017", "gobang", "log")
	if err == nil {
		logger.Hooks.Add(hooker)
	}
}
