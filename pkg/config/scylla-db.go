package config

type scyllaDB struct {
	Hosts    []string
	Database string
	Username string
	Password string
}

var ScyllaDB scyllaDB

var PageSize int = 10
