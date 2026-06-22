package pkg

import "os"

type Config struct {
	AppName    string
	ServerPort string
	DBPath     string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/vempala.db"
	}

	return &Config{
		AppName:    "Vempala Kids",
		ServerPort: port,
		DBPath:     dbPath,
	}
}
