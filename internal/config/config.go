package config

import "github.com/spf13/viper"

// Config stores all configuration for the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DSN string // Data Source Name
}

type JWTConfig struct {
	Secret string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
