package configkit

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Addr     string `mapstructure:"addr" default:":8080"`
	Database string `mapstructure:"database"`
}

var GlobalConfig Config

func Init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config.yaml file not found")
		}
	}

	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println("xx", GlobalConfig.Database)
}
