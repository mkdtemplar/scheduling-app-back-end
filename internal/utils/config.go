package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbSource      string        `mapstructure:"DB_SOURCE"`
	JWTSecret     string        `mapstructure:"JWT_SECRET"`
	Issuer        string        `mapstructure:"JWT_ISSER"`
	Audience      string        `mapstructure:"AUDIENCE"`
	TokenExpiry   time.Duration `mapstructure:"TOKEN_EXPIRY"`
	RefreshExpiry time.Duration `mapstructure:"REFRESH_EXPIRY"`
	CookieDomain  string        `mapstructure:"COOKIE_DOMAIN"`
	CookiePath    string        `mapstructure:"COOKIE_PATH"`
	CookieName    string        `mapstructure:"COOKIE_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return

}
