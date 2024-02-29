package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_PASSWORD string `json:"db_password"`
	DB_USER     string `json:"db_user"`
	DB_URL      string `json:"db_url"`
	HTTPS       bool   `json:"https"`
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	Port        string `json:"port"`
}

func (c *Config) SaveConfig() error {
	f, err := os.Create("./config.json")
	if err != nil {
		return err
	}
	defer f.Close()
	json, _ := json.MarshalIndent(c, "", " ")
	_, err = f.WriteString(string(json))
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("./config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func (c *Config) String() string {
	return fmt.Sprintf(
		`{
	DB_USER: %s
	DB_PASSWORD: %s
	DB_URL: %s
	HTTPS: %t
	PORT: %s
	CERT: %s
	KEY: %s
	}
	`, c.DB_USER, c.DB_PASSWORD, c.DB_URL, c.HTTPS, c.Port, c.Cert, c.Key)
}
