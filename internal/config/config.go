package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"rest-api-service/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_id" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"1234"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once // for singleton pattern

func GetConfig() *Config { // for singleton pattern
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Loading the configuration...")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			description, _ := cleanenv.GetDescription(instance, nil)
			logger.Panic("Error while loading the configuration: ", description)
		}
	})
	return instance
}

func (c *Config) GetListenAddress() string {
	return c.Listen.BindIp + ":" + c.Listen.Port
}
