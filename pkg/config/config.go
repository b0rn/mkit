package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Loads environment variables into the current environment using viper.
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

// Loads a configuration using viper and unmarshals it into a variable.
// Environment variables are expanded into each field.
func BuildConfig(filetype string, filepath string, cfg interface{}) error {
	viper.SetConfigFile(filepath)
	viper.SetConfigType(filetype)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		if v != "" {
			viper.Set(k, os.ExpandEnv(v))
		}
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
