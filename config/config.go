package config

import (
	"github.com/spf13/viper"
)

var config *Config

func GetConfig(test ...bool) *Config {
	if config == nil {
		err := loadConfig(test...)
		if err != nil {
			panic(err)
		}
	}

	return config
}

type Config struct {
	App App `yaml:"app"`
}

type App struct {
	ImageDirPath string `yaml:"imageDirPath"`
}

func loadConfig(test ...bool) error {
	viperInstance := viper.New()
	var name string
	if test != nil && test[0] {
		name = "config_test"
	} else {
		name = "config"
	}
	viperInstance.SetConfigName(name)
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath("./config/")
	viperInstance.AddConfigPath("../config/")
	viperInstance.AutomaticEnv()
	err := viperInstance.ReadInConfig()
	if err != nil {
		return err
	}

	config = &Config{}
	if err := viperInstance.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
