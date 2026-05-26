package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config contiene toda la configuración tipada de la aplicación.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port int
	Env  string // development | production
}

type DatabaseConfig struct {
	URL string // postgres://user:pass@host:port/dbname
}

type RedisConfig struct {
	Addr     string // host:port
	Password string
	DB       int
}

type JWTConfig struct {
	Secret     string
	Expiration int // horas
}

// Cargar lee .env y devuelve una Config poblada.
func Cargar() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[CONFIG] .env no encontrado, usando variables de entorno del sistema")
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: obtenerInt("SERVER_PORT", 8080),
			Env:  obtenerStr("ENV", "development"),
		},
		Database: DatabaseConfig{
			URL: obtenerStr("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/deprodeca?sslmode=disable"),
		},
		Redis: RedisConfig{
			Addr:     obtenerStr("REDIS_ADDR", "localhost:6379"),
			Password: obtenerStr("REDIS_PASSWORD", ""),
			DB:       obtenerInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret:     obtenerStr("JWT_SECRET", "deprodeca-secret-dev"),
			Expiration: obtenerInt("JWT_EXPIRATION_HOURS", 24),
		},
	}

	log.Printf("[CONFIG] Configuración cargada — puerto=%d, env=%s", cfg.Server.Port, cfg.Server.Env)
	return cfg
}

func obtenerStr(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func obtenerInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return defaultVal
}
