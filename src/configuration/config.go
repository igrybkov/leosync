package configuration

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type linguaLeo struct {
	Email    string
	Password string
}

type Config struct {
	LinguaLeo linguaLeo
}

func GetConfig() Config {
	config := Config{
		LinguaLeo: linguaLeo{
			Email:    cast.ToString(viper.Get("lingualeo.email")),
			Password: cast.ToString(viper.Get("lingualeo.password")),
		},
	}
	return config
}
