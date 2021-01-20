package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Configer *viper.Viper

func init() {
	Configer = viper.New()
	Configer.SetConfigName("config")
	Configer.SetConfigType("yaml")
	Configer.AddConfigPath(".")
	Configer.AutomaticEnv()

	err := Configer.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
