package pkg

type Config struct {
	AppName    string
	ServerPort string
	DBPath     string
}

func LoadConfig() *Config {
	return &Config{
		AppName:    "Vempala Kids",
		ServerPort: "8080",
		DBPath:     "data/vempala.db",
	}
}
