package config

type Config struct {
	DbConfig *DbConfig
}

func NewConfig() *Config {
	return &Config{
		DbConfig: buildDbconfig(),
	}
}

func buildDbconfig() *DbConfig {
	return &DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "wallets",
	}
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}
