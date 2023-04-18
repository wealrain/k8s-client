package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port struct {
			Api string `yaml:"api"`
		} `yaml:"port"`
	} `yaml:"server"`

	Redis struct {
		Addr string `yaml:"addr"`
	} `yaml:"redis"`

	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"mysql"`
}

var (
	config      *Config
	configMutex sync.Mutex
)

func LoadConfig() *Config {
	if config != nil {
		return config
	}

	configMutex.Lock()
	defer configMutex.Unlock()

	// double check
	if config != nil {
		return config
	}

	file, err := os.Open("./config.yaml")

	if err != nil {
		log.Fatalln("can not open config file", err)
	}
	defer file.Close()

	config := Config{}
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		log.Fatalln("parse config file error", err)
	}

	return &config
}
