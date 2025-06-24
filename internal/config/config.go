package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env           string `yaml:"env"`
	FileName      string `yaml:"file_name"`
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
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	CollName string `yaml:"coll_name"`
}

func Load() *Config {
	cfg := new(Config)

	if err := cleanenv.ReadConfig("config.yaml", cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
