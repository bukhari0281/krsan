package models

type Config struct {
	ServerPort string
	Database   Database
}

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
