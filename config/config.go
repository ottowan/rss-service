package config

import (
	"log"

	"github.com/spf13/viper"
)

// Configuration struct
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Gateway  GatewayConfiguration
}

// InitConfig config
func InitConfig(mode []string) Configuration {
	var configuration Configuration

	if len(mode) == 1 {
		viper.SetConfigName("config_local")
	} else {
		switch mode[1] {
		case "uat":
			viper.SetConfigName("config_uat")
		case "product":
			viper.SetConfigName("config_production")
		default:
			viper.SetConfigName("config_local")
		}

	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return configuration
}
