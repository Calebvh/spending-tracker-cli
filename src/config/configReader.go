package config

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

type Config struct {
	TransactionFolder string
}

func ReadConfig() (*Config, err) {
	//print out current directory
	dir, err := os.Getwd()
	println(dir)

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Failed to load config.ini:", err)
		return err
	}

	// Create a Config struct to store the configuration values
	config := Config{}

	// Read the values from the configuration file
	err = cfg.Section("Transactions").MapTo(&config)
	if err != nil {
		log.Fatal("Failed to read transaction configuration:", err)
	}

	return config
}
