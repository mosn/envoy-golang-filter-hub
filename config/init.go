package config

import (
	"envoy-golang-filter-hub/utils"
	"github.com/creasty/defaults"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"

	"os"
)

var configPath = "config/config.yaml"
var c = new(GlobalConfig)

func Init() {
	if err := defaults.Set(c); err != nil {
		panic(err)
	}

	if !utils.FileExist(configPath) {
		log.Println("Config file not exist")
		gen()
		os.Exit(0)
	}
	load()
	validate()
}

func load() {
	viper.SetConfigFile(configPath)
	utils.PanicIfErr(
		viper.ReadInConfig(),
		viper.Unmarshal(c),
	)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(c); err != nil {
			log.Println(err)
		}
	})

}

func gen() {
	log.Println("Generating config file in ", configPath, "...")
	bytes, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	if err := utils.FileCreate(configPath, bytes); err != nil {
		panic(err)
	}
	log.Println("Generate config file success")
}

func Get() GlobalConfig {
	return *c
}

func validate() {
	if err := validator.New().Struct(c); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
