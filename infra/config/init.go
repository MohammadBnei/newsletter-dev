package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var cfg *Config

// GetConfig returns the configuration for the application.
func GetConfig() Config {
	if cfg != nil {
		return *cfg
	}

	Init()
	return *cfg
}

func Init() *Config {
	viper.SetConfigName("MONGO_SECRET_FILE")
	viper.SetConfigType("env")

	viper.AddConfigPath("./")
	viper.AddConfigPath("/etc/secrets")

	viper.AutomaticEnv()

	// Read the configuration
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read .env file : %s", err)
	}

	cfg = &Config{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal .env file", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		Init()
	})

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(cfg); err != nil {
		log.Fatalf("Failed to validate config : %s", err)
	}

	return cfg

}
