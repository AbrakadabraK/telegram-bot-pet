package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	filename = `D:/project/telegram-zero-start-proj/telegram-bot-pet/config/config.yaml`
)

type TelegramConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Config struct {
	Telegram []TelegramConfig `yaml:"telegram"`
}

// LoadConfig загружает конфигурацию из файла YAML
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetToken(cfg *Config, nameKey string) (string, error) {
	var res string
	for _, t := range cfg.Telegram {
		if t.Name == nameKey {
			res = t.Value
			break
		}
	}

	if res != "" {
		return res, nil
	}
	return res, errors.New("not found value by key")
}
