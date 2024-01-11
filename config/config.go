package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

var configPath = "config.yaml"

type Config struct {
	HTTP struct {
		// http or https
		Protocol     string        `yaml:"protocol"`
		Host         string        `yaml:"host"`
		Port         int           `yaml:"port"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
		SertFile     string        `yaml:"sert_file"`
		SertKey      string        `yaml:"sert_key"`
	} `yaml:"http"`
	Postgres Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Settings struct {
		MaxOpenConns    int           `yaml:"max_open_conns"`
		ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time"`
		MaxIdleConns    int           `yaml:"max_idle_conns"`
		MaxIdleLifeTime time.Duration `yaml:"max_idle_life_time"`
	}
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	cfg := Config{}

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath != "" {
		configPath = cfgPath
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
