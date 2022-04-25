package internal

import (
	"github.com/gLRB/gGateway/internal/backends"
	"github.com/spf13/viper"
)

type Env int

const (
	EnvProd Env = iota
	EnvDev
)

func (e Env) String() string {
	if e == EnvProd {
		return "prod"
	} else if e == EnvDev {
		return "dev"
	}
	return "unknown"
}

// TODO
var env Env = EnvDev

func Init() {
	initConfig()
	backends.Init()
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./etc/conf")

	viper.SetConfigName("config-default")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config-" + env.String())
	err = viper.MergeInConfig()
	if err != nil {
		panic(err)
	}
}
