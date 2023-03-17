package config

import "github.com/spf13/viper"

type Config struct {
	SvConfig *SvConfig
	DbConfig *DbConfig
}

type SvConfig struct {
	Port string
	JWT  JWTConfig
}

type JWTConfig struct {
	Key []byte
}

type DbConfig struct {
	User        string
	Password    string
	Name        string
	Instance    string
	Credentials string
}

func LoadConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var svConfig = SvConfig{
		Port: viper.GetString("server.port"),
		JWT: JWTConfig{
			Key: []byte(viper.GetString("server.JWT.key")),
		},
	}
	var dbConfig = DbConfig{
		User:        viper.GetString("database.user"),
		Password:    viper.GetString("database.password"),
		Name:        viper.GetString("database.name"),
		Instance:    viper.GetString("database.instance"),
		Credentials: viper.GetString("database.credentials"),
	}

	return &Config{
		SvConfig: &svConfig,
		DbConfig: &dbConfig,
	}, nil
}
