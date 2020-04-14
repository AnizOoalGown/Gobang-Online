package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Config *viper.Viper

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}

	Config = viper.New()
	Config.AddConfigPath(path)
	Config.SetConfigName("config")
	Config.SetConfigType("yml")

	if err = Config.ReadInConfig(); err != nil {
		log.Error(err)
	}
}
