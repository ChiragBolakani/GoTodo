package db

import (
	"go_tutorials/pkg/config"

	"github.com/gocql/gocql"
)

type DB struct {
	client  *gocql.ClusterConfig
	session *gocql.Session
}

func NewDB() *DB {
	client := gocql.NewCluster(config.ScyllaDB.Hosts...)

	client.Authenticator = gocql.PasswordAuthenticator{
		Username: config.ScyllaDB.Username,
		Password: config.ScyllaDB.Password,
	}
	client.Keyspace = config.ScyllaDB.Database

	client.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	return &DB{
		client: client,
	}
}

func (d *DB) GetConn() *gocql.Session {
	if d.session == nil {
		session, err := d.client.CreateSession()
		if err != nil {
			panic(err)
		}
		d.session = session
	}

	return d.session
}
