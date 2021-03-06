package store

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// initConfig will setup the config reader
func initConfig(cfg Config) {
	// Register alias to support mapping of env with config
	viper.RegisterAlias("authentication.username", "HEAMON_USER")
	viper.RegisterAlias("authentication.password", "HEAMON_PASS")
	viper.RegisterAlias("monitor.interval", "HEAMON_INTERVAL")

	// Set defaults
	viper.SetDefault("TITLE", "Heamon")
	viper.SetDefault("PORT", "5000")
	viper.SetDefault("HEAMON_USER", "admin")
	viper.SetDefault("HEAMON_PASS", "password")
	viper.SetDefault("HEAMON_INTERVAL", "60")

	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/heamon/")
	viper.AddConfigPath("$HOME/.heamon")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.Fatal("failed to read config:", err)
		}
	}

	replacePlaceholders()

	viper.Unmarshal(&cfg)
}

func replacePlaceholders() {
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${{.") && strings.HasSuffix(value, "}}") {
			viper.Set(k, getEnv(strings.TrimSuffix(strings.TrimPrefix(value, "${{."), "}}")))
		}
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logrus.Warnf("environmental variable %s is \"\"", key)
	}

	return val
}
