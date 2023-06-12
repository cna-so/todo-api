package helpers

import "fmt"

type DatabaseConfig struct {
	Host     string
	User     string
	Port     string
	SSLMode  string
	Database string
	Password string
}

func (cfg *DatabaseConfig) DbValues() string {
	return fmt.Sprintf("user=%s port=%s host=%s sslmode=%s dbname=%s password=%s", cfg.User, cfg.Port, cfg.Host, cfg.SSLMode, cfg.Database, cfg.Password)
}
