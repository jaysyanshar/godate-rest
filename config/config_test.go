package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigLoading(t *testing.T) {
	// Create a temporary .env file with test values
	envContent := "PORT=8080"
	envFile := ".env.test"
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(envFile)

	// Set Viper to read the test .env file
	viper.SetConfigFile(envFile)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Error reading .env file: %v", err)
	}

	// Initialize the config
	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		t.Fatalf("Error unmarshalling config: %v", err)
	}

	// Assert the values are correctly loaded
	assert.Equal(t, "8080", config.Port)
}
