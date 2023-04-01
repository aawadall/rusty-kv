package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Configuration Manager

type ConfigurationManager struct {
	Configuration map[string]interface{}
	logger        *log.Logger
}

// NewConfigurationManager - A function that creates a new Configuration Manager
func NewConfigurationManager(configMap map[string]string) *ConfigurationManager {
	cfg := &ConfigurationManager{
		Configuration: make(map[string]interface{}),
		logger:        log.New(os.Stdout, "CONFIG: ", log.LstdFlags),
	}
	// append configuration from configMap
	for key, value := range configMap {
		cfg.Configuration[key] = value
	}

	cfg.logger.Println("Configuration Manager created")
	cfg.LoadFromEnvironment()
	return cfg
}

// Get - A function that gets a configuration value
func (c *ConfigurationManager) Get(key string) (value interface{}, err error) {
	// if the key is empty return error
	if key == "" {
		return nil, fmt.Errorf("configuration `Key` cannot be empty")
	}

	// if the configuration is nil, return error
	if c.Configuration == nil {
		return nil, fmt.Errorf("configuration is nil")
	}

	// if the key is not in the configuration, return error
	if _, ok := c.Configuration[key]; !ok {
		return nil, fmt.Errorf("configuration `Key` not found")
	}

	return c.Configuration[key], nil
}

// Set - A function that sets a configuration value
func (c *ConfigurationManager) Set(key string, value interface{}) (err error) {
	// if the key is empty return error
	if key == "" {
		return fmt.Errorf("configuration `Key` cannot be empty")
	}

	// if the value is empty return error
	if value == nil {
		return fmt.Errorf("configuration `Value` cannot be empty")
	}

	// if the configuration is nil, create it
	if c.Configuration == nil {
		c.Configuration = make(map[string]interface{})
	}

	c.Configuration[key] = value

	return nil
}

// LoadFromEnvironment - A function that loads the configuration from the environment
func (c *ConfigurationManager) LoadFromEnvironment() (err error) {
	c.logger.Println("Loading configuration from environment")
	// Get all environment variables
	envVars := os.Environ()

	// Iterate over the environment variables
	for _, envVar := range envVars {
		// Split the environment variable into key and value
		key, value := splitEnvVar(envVar)
		// Set the configuration
		c.Set(key, value)
	}

	return nil
}

// GetConfig - A function that gets all configuration
func (c *ConfigurationManager) GetConfig() map[string]interface{} {
	return c.Configuration
}

// Helper Functions
// splitEnvVar - A function that splits an environment variable into key and value
func splitEnvVar(envVar string) (key string, value string) {
	// split string by =
	split := strings.Split(envVar, "=")
	key = split[0]

	// value is the concatenation of the rest of the split string with the = inserted back
	value = strings.Join(split[1:], "=")

	return key, value
}
