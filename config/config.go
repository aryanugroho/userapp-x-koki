package config

type Config struct {
	DB DB
}

type DB struct {
	Host     string
	Name     string
	UserName string
	Password string
	Port     string
}
