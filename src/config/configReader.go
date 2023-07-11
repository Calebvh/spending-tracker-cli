package config

import (
	"log"

	"github.com/go-ini/ini"
)

type Config struct {
	TransactionFolder string
}

func ReadConfig() *Config {
	//print out current directory

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Failed to load config.ini:", err)
		return nil
	}

	// Create a Config struct to store the configuration values
	config := &Config{}

	// Read the values from the configuration file
	err = cfg.Section("Transactions").MapTo(&config)
	if err != nil {
		log.Fatal("Failed to read transaction configuration:", err)
		return nil
	}

	//Verify Config struct is populated
	if config.TransactionFolder == "" {
		log.Fatal("Transaction folder not set in config.ini")
		return nil
	}

	return config
}
