package configs

import (
	sharedConfigs "project1.v0/configs"
)

type Config struct {
	Db sharedConfigs.DbConfig
}

func LoadConfig() *Config {
	return &Config{
		Db: sharedConfigs.NewDbConfig(),
	}
}
