package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env           string `yaml:"env"`
	URLs          string `yaml:"urls"`
	ConfigServer  `yaml:"server"`
	ConfigMongoDB `yaml:"mongodb"`
}

type ConfigServer struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type ConfigMongoDB struct {
	Username       string        `yaml:"username"`
	Password       string        `yaml:"password"`
	Host           string        `yaml:"host"`
	Port           string        `yaml:"port"`
	DBName         string        `yaml:"db_name"`
	CollName       string        `yaml:"coll_name"`
	ConnectTimeout time.Duration `yaml:"conn_timeout"`
	SelectTimeout  time.Duration `yaml:"select_timeout"`
}

func MustLoad(path string) *Config {
	cfg := new(Config)

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return cfg
}
