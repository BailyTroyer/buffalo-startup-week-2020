package util

import "github.com/spf13/viper"

// Config API config
type Config struct {
	Meta     MetaConfig
	NewRelic NewRelicConfig
}

// MetaConfig HTTP API config
type MetaConfig struct {
	Port int
	Name string
}

// NewRelicConfig NewRelic API client config
type NewRelicConfig struct {
	Token string
}

// LoadConfig load config.yaml and unmarshall into config struct
func LoadConfig() (Config, error) {

	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath("/secrets")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
