package config

type Config struct {
	MongoURI string `mapstructure:"MONGO_URI" validate:"required,uri"`
	MongoDB  string `mapstructure:"MONGO_DB" validate:"required"`

	Port string `mapstructure:"PORT"`

	Local bool
}
