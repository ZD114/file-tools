package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	DefaultConfigFile = "application.yml"
)

type SpringConfiguration struct {
	Server ServerProperties `yaml:"server"`
}

type ServerProperties struct {
	Port int32 `yaml:"port"`
}

var springConfiguration SpringConfiguration

func init() {
	conf := new(SpringConfiguration)

	f := DefaultConfigFile

	if data, err := os.ReadFile(f); err != nil {
		return
	} else if err = yaml.Unmarshal(data, &conf); err != nil {
		return
	}

	springConfiguration = *conf
}

func GetConfig() SpringConfiguration {
	return springConfiguration
}
