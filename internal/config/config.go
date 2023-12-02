package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	DB struct {
		User     string `yaml:"postgres_user" env-default:"postgres"`
		Password string `yaml:"postgres_password" env-default:"postgres"`
		DBName   string `yaml:"postgres_dbname" env-default:"postgres"`
		Host     string `yaml:"postgres_host" env-default:"database"`
		Port     string `yaml:"postgres_port" env-default:"5432"`
		SSLMode  string `yaml:"postgres_sslmode" env-default:"disable"`
	} `yaml:"db"`
	Stan struct {
		ClusterID     string `yaml:"stan_cluster_id" env-default:"wb_stan_cluster"`
		ClientID      string `yaml:"stan_client_id" env-default:"wb_stan_client"`
		URL           string `yaml:"stan_url" env-default:"nats://stan:4222"`
		Subject       string `yaml:"stan_subject" env-default:"wb_stan_subject"`
		Durable       string `yaml:"stan_durable" env-default:"wb_stan_durable"`
		StartDeltaMin int    `yaml:"start_delta_min" env-default:"10"`
	} `yaml:"stan"`
	App struct {
		Port string `yaml:"app_port" env-default:":8080"`
	} `yaml:"app"`
}

func Parse() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, errors.Wrap(err, "parse config error")
	}

	return &cfg, nil
}

func ParseWithFile(path string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, errors.Wrap(err, "parse config error")
	}

	return &cfg, nil
}
