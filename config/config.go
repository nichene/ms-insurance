package config

import "github.com/spf13/viper"

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

func LoadEnvVars(opts ...Options) *Config {
	options := configOptions{}

	for _, o := range opts {
		o.apply(&options)
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AddConfigPath(options.configPath)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// Defaults
	viper.SetDefault("PORT", "8080")

	return &Config{
		Port:   viper.GetString("PORT"),
		DBHost: viper.GetString("DB_HOST"),
		DBName: viper.GetString("DB_NAME"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASS"),
	}
}

type configOptions struct {
	configPath string
}

type Options interface {
	apply(*configOptions)
}

type configPathOption string

func (c configPathOption) apply(opts *configOptions) {
	opts.configPath = string(c)
}

func WithConfigPath(configPath string) Options {
	return configPathOption(configPath)
}
