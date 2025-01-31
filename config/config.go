package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	FrontendURL string `mapstructure:"FRONTEND_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      int    `mapstructure:"DB_PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
