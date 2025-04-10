package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	API    APIConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Environment  string
}

type APIConfig struct {
	RateLimit     int
	RateLimitTime time.Duration
}

func Load() (*Config, error) {
	godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
			Environment:  getEnv("ENVIRONMENT", "development"),
		},
		API: APIConfig{
			RateLimit:     getIntEnv("RATE_LIMIT", 50),
			RateLimitTime: getDurationEnv("RATE_LIMIT_TIME", time.Minute),
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.Server.Environment == "production"
}
