package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8845")
	viper.SetDefault("THEME", "cerulean")
	viper.SetDefault("KEYPATH", "/data/gans/ssh")
	// viper.SetDefault("KEYPATH", "/root/.ssh")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.KeyPath, _ = viper.Get("KEYPATH").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("keypath", config.KeyPath)

	err := viper.WriteConfig()
	check.IfError(err)
}
