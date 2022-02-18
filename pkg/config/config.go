package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server struct {
		Host string `validate:"required" yaml:"host"`
		Port string `validate:"required" yaml:"port"`
	} `validate:"required" yaml:"server"`
	Database struct {
		Host string `validate:"required" yaml:"host"`
		Port string `validate:"required" yaml:"port"`
		Name string `validate:"required" yaml:"name"`
		User string `validate:"required" yaml:"user"`
		Ssl  string `validate:"required" yaml:"ssl"`
	} `validate:"required" yaml:"database"`
	Jwt struct {
		Secret string `validate:"required" yaml:"secret"`
	}
}

var config Configuration

func ReadConf() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return err
	}

	return nil
}
