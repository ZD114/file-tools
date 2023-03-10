package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"zhangda/file-tools/util"
)

const (
	DefaultConfigFile = "application.yml"
)

type SpringConfiguration struct {
	Server  ServerProperties  `yaml:"server"`
	Logging LoggingProperties `yaml:"logging"`
}

type ServerProperties struct {
	Port int32 `yaml:"port"`
}

type LoggingProperties struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

var springConfiguration SpringConfiguration

func init() {
	conf := new(SpringConfiguration)

	f := DefaultConfigFile

	if data, err := os.ReadFile(f); err != nil {
		util.Logger.Error("config", util.Any("config", err))

		return
	} else if err = yaml.Unmarshal(data, &conf); err != nil {
		util.Logger.Error("config", util.Any("config", err))

		return
	}

	springConfiguration = *conf
}

func GetConfig() SpringConfiguration {
	return springConfiguration
}
