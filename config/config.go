package config

type Config struct {
	DbConfig *DbConfig
	ServerConfig *ServerConfig
}

func NewConfig() *Config {
	return &Config{
		DbConfig: buildDbconfig(),
		ServerConfig: buildServerConfig(),
	}
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type ServerConfig struct {
	Host string
	Port string
}

func buildDbconfig() *DbConfig {
	return &DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",
	}
}

func buildServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: "localhost",
		Port: "8080",
	}
}
