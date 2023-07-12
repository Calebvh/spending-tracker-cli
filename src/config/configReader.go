package config

import (
	"log"

	"github.com/go-ini/ini"
)

type ConfigSection string

const (
	TRANSACTIONS ConfigSection = "Transactions"
	LOGGER       ConfigSection = "Logger"
)

type TransactionConfig struct {
	TransactionFolder string
}

type LoggerConfig struct {
	LogLevel    string
	DebugMode   bool
	WarningMode bool
}

func ReadTransactionConfig() *TransactionConfig {
	//print out current directory

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Failed to load config.ini:", err)
		return nil
	}

	// Create a Config struct to store the configuration values
	config := &TransactionConfig{}

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

func ReadLoggerConfig() *LoggerConfig {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Failed to load config.ini:", err)
		return nil
	}

	// Create a Config struct to store the configuration values
	config := &LoggerConfig{}

	// Read the values from the configuration file
	err = cfg.Section("Logger").MapTo(&config)
	if err != nil {
		log.Fatal("Failed to read logger configuration:", err)
		return nil
	}

	//Verify Config struct is populated
	if config.LogLevel == "" {
		log.Fatal("Log level not set in config.ini")
		return nil
	}

	return config
}
