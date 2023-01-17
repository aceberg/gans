package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/gans/internal/check"
	"github.com/aceberg/gans/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("DB", "/data/gans/sqlite.db")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8845")
	viper.SetDefault("THEME", "cerulean")
	viper.SetDefault("SHOW", "25")
	viper.SetDefault("YAMLPATH", "/data/gans/repo.yaml")
	viper.SetDefault("KEYPATH", "/data/gans/ssh")
	viper.SetDefault("INTERVAL", "5s")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DB, _ = viper.Get("DB").(string)
	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Show, _ = viper.Get("SHOW").(string)
	config.KeyPath, _ = viper.Get("KEYPATH").(string)
	config.YamlPath, _ = viper.Get("YAMLPATH").(string)
	config.Interval, _ = viper.Get("INTERVAL").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("db", config.DB)
	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("show", config.Show)
	viper.Set("keypath", config.KeyPath)
	viper.Set("yamlpath", config.YamlPath)
	viper.Set("interval", config.Interval)

	err := viper.WriteConfig()
	check.IfError(err)
}
