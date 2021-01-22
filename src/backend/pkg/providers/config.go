package providers

import "github.com/spf13/viper"

type Config struct {
	DBPort     string
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
}

func LoadConfig() *Config {
	// set default
	config := &Config{
		DBPort:     "5440",
		DBName:     "postgres",
		DBUser:     "root",
		DBPassword: "secret",
		DBHost:     "localhost",
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil
	}
	return config
}
