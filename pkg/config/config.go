package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadEnvVars(filetype string, filepath string) error {
	viper.SetConfigFile(filepath)
	viper.SetConfigType(filetype)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for v, i := range viper.AllSettings() {
		os.Setenv(strings.ToUpper(v), i.(string))
	}
	return nil
}

func BuildConfig(filetype string, filepath string, cfg interface{}) error {
	viper.SetConfigFile(filepath)
	viper.SetConfigType(filetype)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
