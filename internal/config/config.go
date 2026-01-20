package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"ENV" env-default:"development"`

	DatabaseHost     string `env:"DB_HOST" env-required:"true"`
	DatabaseName     string `env:"DB_NAME" env-required:"true"`
	DatabaseUser     string `env:"DB_USER" env-required:"true"`
	DatabasePassword string `env:"DB_PASSWORD" env-required:"true"`
	DatabasePort     string `env:"DB_PORT" env-required:"true"`

	RabbitHost     string `env:"RABBIT_HOST" env-required:"true"`
	RabbitPort     string `env:"RABBIT_PORT" env-required:"true"`
	RabbitUser     string `env:"RABBIT_USER" env-required:"true"`
	RabbitPassword string `env:"RABBIT_PASSWORD" env-required:"true"`
}

func MustLoad() *Config {
	configPath := "./.env"
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	// Проверяем существование конфиг-файла
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	// Читаем конфиг-файл и заполняем нашу структуру
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
