package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	pg_host     = "localhost"
	pg_port     = "5432"
	pg_user     = "postgres"
	pg_password = "postgres"
	pg_name     = "postgres"
	pg_sslmode  = "disable"
	http_port   = "8080"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("config - No .env file: %v", err)
	}
}

type (
	Config struct {
		HTTP
		PG
	}

	HTTP struct {
		Port string
	}
	PG struct {
		Host     string
		Port     string
		User     string
		Password string
		DBname   string
		SSLmode  string
	}
)

func NewConfig() (*Config, error) {
	return &Config{*newHTTP(), *newPG()}, nil

}
func newPG() *PG {
	return &PG{
		Host:     getEnv("PG_HOST", pg_host),
		Port:     getEnv("PG_PORT", pg_port),
		User:     getEnv("PG_USER", pg_user),
		Password: getEnv("PG_PASSWORD", pg_password),
		DBname:   getEnv("PG_NAME", pg_name),
		SSLmode:  getEnv("PG_SSLmode", pg_sslmode),
	}
}
func newHTTP() *HTTP {
	return &HTTP{
		Port: getEnv("HTTP_PORT", http_port),
	}
}
func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue // если нет переменной с именем key
}
func (c PG) String() string {
	return fmt.Sprintf("host = %s port=%s user=%s password = %s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBname, c.SSLmode)
}
