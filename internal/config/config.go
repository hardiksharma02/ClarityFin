package config

import "github.com/spf13/viper"

// Config stores all configuration for the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	SMS      SMSConfig
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

type SMSConfig struct {
	Provider string
	Twilio   TwilioConfig
	MSG91    MSG91Config
}

type TwilioConfig struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

type MSG91Config struct {
	APIKey   string
	SenderID string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
