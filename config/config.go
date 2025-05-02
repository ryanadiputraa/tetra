package config

import "github.com/spf13/viper"

type Config struct {
	Port                string `mapstructure:"PORT"`
	FrontendURL         string `mapstructure:"FRONTEND_URL"`
	JWTSecret           string `mapstructure:"JWT_SECRET"`
	EncryptionKey       string `mapstructure:"ENCRYPTION_KEY"`
	DashboardServiceURI string `mapstructure:"DASHBOARD_SERVICE_URI"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	DBPort              int    `mapstructure:"DB_PORT"`
	RedisAddr           string `mapstructure:"REDIS_ADDR"`
	RedisPassword       string `mapstructure:"REDIS_PASSWORD"`
	OauthState          string `mapstructure:"OAUTH_STATE"`
	OauthRedirectURL    string `mapstructure:"OAUTH_REDIRECT_URL"`
	GoogleClientID      string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret  string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	SMTPEmail           string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword        string `mapstructure:"SMTP_PASSWORD"`
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
