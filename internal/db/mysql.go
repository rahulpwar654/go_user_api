package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func LoadConfig() Config {
	return Config{
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "3306"),
		Name: getEnv("DB_NAME", "go-userdb"),
		User: getEnv("DB_USER", "root"),
		Pass: getEnv("DB_PASS", "root"),
	}
}

func OpenMySQL(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
