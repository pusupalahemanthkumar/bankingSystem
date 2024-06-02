package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// Util Function to load configs for from environment file or environment varaibles
func LoadConfig(path string) (config Config, err error) {
	// path like `../ or ./`
	viper.AddConfigPath(path)
	// FileName
	viper.SetConfigName("app")
	// Type like `env or json or xml`
	viper.SetConfigType("env")

	// load environment variables from runtime cloud (aws, azure, gcp, local)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return

}
