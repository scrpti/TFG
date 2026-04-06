// puerto de la app
// host de la base de datos
// usuario
// contraseña
// nombre de la BD
// carpeta de uploads


// Uso del paquete 
// Paquete config que se encarga de contener los parametros de configuracion para arrancar la app

package config

import(
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DB DBConfig
	Storage StorageConfig
}

type DBConfig struct{
	Host string
	Port string
	User string
	Password string
	Name string
	SSLMode string
}

type StorageConfig struct{
	UploadDir string
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("enviroment variable %s is required", key)
	}
	return value, nil
}

func Load() (*Config, error) {
	//Cargar .env

	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file found")
	}

	//Leer variables obligatorias

	appPort, err := getEnv("APP_PORT")
	if err != nil {
		return nil, err
	}

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return nil, err
	}

	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	
	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	dbSSLMode, err := getEnv("DB_SSLMODE")
	if err != nil {
		return nil, err
	}

	uploadDir, err := getEnv("UPLOAD_DIR")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AppPort: appPort,
		DB : DBConfig{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Password: dbPassword,
			Name: dbName,
			SSLMode: dbSSLMode,
		},
		Storage: StorageConfig{
			UploadDir: uploadDir,
		},
	}

	return cfg, nil
}