package config

import (
	"github.com/spf13/viper"
)

var config *Config

func GetConfig() *Config {
	if config == nil {
		err := loadConfig()
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

func loadConfig() error {
	viperInstance := viper.New()
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath("./config/")
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
