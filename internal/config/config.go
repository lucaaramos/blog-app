package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func LoadConfig() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Asignar valores a las variables de configuración
	AppConfig = Config{
		Port:   getEnv("PORT", "8000"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "27017"),
		// DBUser: getEnv("DB_USER", "root"),
		// DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName: getEnv("DB_NAME", "blog"),
	}
}

// Helper function to get environment variables or default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
