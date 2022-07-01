package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var (
	instance *configManager
)

type configManager struct {
	vip *viper.Viper
}

func Instance() *configManager {
	if instance == nil {
		vip := viper.New()
		vip.SetConfigName("config")
		vip.AddConfigPath(".")
		vip.AddConfigPath("./config")

		if err := vip.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Info("Config not found, load default configs")
				setDefaults(vip)
			} else {
				log.Info("Error on read config file #{err}")
			}
		}
		instance = &configManager{vip: vip}
	}
	return instance
}

func setDefaults(vip *viper.Viper) {
	vip.SetDefault("server.host", "0.0.0.0:8080")
	vip.SetDefault("app.name", "package")
	vip.SetDefault("app.logLevel", "info")
}

func (cm *configManager) GetAppConfig() AppConfig {
	vipAppConfig := cm.vip.Sub("app")
	return AppConfig{
		Name:     vipAppConfig.GetString("name"),
		LogLevel: vipAppConfig.GetString("logLevel"),
	}
}

func (cm *configManager) GetServerConfig() ServerConfig {
	vipServerConfig := cm.vip.Sub("server")
	return ServerConfig{
		Host: vipServerConfig.GetString("host"),
	}
}
