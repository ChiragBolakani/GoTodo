package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()

	ScyllaDB = scyllaDB{
		Hosts:    strings.Split(os.Getenv("SCYLLA_DB_HOSTS"), ","),
		Database: os.Getenv("SCYLLA_DB_DATABASE"),
		Username: os.Getenv("SCYLLA_DB_USERNAME"),
		Password: os.Getenv("SCYLLA_DB_PASSWORD"),
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 3000
	}

	Server = server{
		Port: port,
	}
}
