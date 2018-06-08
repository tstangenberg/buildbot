package buildbot

import (
	"github.com/spf13/viper"
	"github.com/labstack/gommon/log"
)

type Configuration struct {
	Task CmdTask `yaml:"task"`
}

type CmdTask struct {
	Command string `yaml:"command"`
	Arguments []string `yaml:"arguments"`
}

func init() {
	viper.SetConfigName("buildbot.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func (c *Configuration) Initialize() {
	if err := viper.ReadInConfig(); err != nil {
		if _ , ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfig()
		} else {
			panic(err)
		}
	}
	if err := viper.Unmarshal(c); err != nil {
		panic(err)
	} else {
		log.Info(c)
	}
}


