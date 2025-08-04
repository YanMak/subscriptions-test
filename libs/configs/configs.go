package sharedConfigs

import (
	"fmt"
	"os"
	"time"

	"project1.v0/mappers/conv"
)

type RunModeType string

const (
	RUN_MODE_TYPE_LOCAL RunModeType = "local"
	RUN_MODE_TYPE_K8S   RunModeType = "k8s"
	RUN_MODE_TYPE_CI    RunModeType = "ci"
)

type EnvironmentModeType string

const (
	ENV_MODE_DEVELOPMENT EnvironmentModeType = "development"
	RUN_MODE_STAGING     EnvironmentModeType = "staging"
	RUN_MODE_PRODUCTION  EnvironmentModeType = "production"
)

// ///////
// DB
type DbConfig struct {
	Dsn             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

func NewDbConfig() DbConfig {
	key := "CONN_MAX_IDLE_TIME"
	connMaxIdleTime, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		println("reading .env variable (", key, ") : ", err)
		panic(err)
	}

	return DbConfig{
		//Dsn:             os.Getenv("DSN"),
		Dsn: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
		ConnMaxIdleTime: connMaxIdleTime,
		MaxOpenConns:    conv.ToInt("MAX_OPEN_CONS"),
		MaxIdleConns:    conv.ToInt("MAX_IDLE_CONS"),
	}
}
