package config

import (
	"envoy-go-fliter-hub/tools"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var Path = "config/config.yaml"
var Config = config{}

func Init() {
	viper.SetConfigFile(Path)

	if tools.FileExist(Path) {
		tools.PanicIfErr(
			viper.ReadInConfig(),
			viper.Unmarshal(&Config),
		)
	} else {
		fmt.Println("Config file not exist in ", Path, ". Using environment variables.")
		if err := envconfig.Process("", &Config); err != nil {
			panic(err)
		}
	}

	//fmt.Printf("%+v\n", Config)
}
