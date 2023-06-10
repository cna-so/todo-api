package handlers

import "fmt"

type DatabaseConfig struct {
	User     string
	Port     string
	SSLMode  string
	Database string
	Password string
}

func (cfg *DatabaseConfig) DbValues() string {
	return fmt.Sprintf("user=%s port=%s sslmode=%s dbname=%s password=%s", cfg.User, cfg.Port, cfg.SSLMode, cfg.Database, cfg.Password)
}
