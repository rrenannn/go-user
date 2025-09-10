package config

import (
	"database/sql"
	"log"

	db "github.com/rrenannn/go-user/db/sqlc"
)

type ContainerDI struct {
	Config Config
	Conn   *sql.DB
}

func NewDB(container *ContainerDI) *sql.DB {
	return container.Conn
}

func NewQueries(container *ContainerDI) *db.Queries {
	return db.New(container.Conn)
}

func NewContainerDI(config Config) *ContainerDI {
	c := &ContainerDI{Config: config}
	c.db()
	return c
}

func (c *ContainerDI) db() {
	dbConfig := ConfigDB{
		Host:        c.Config.DBHost,
		Port:        c.Config.DBPort,
		User:        c.Config.DBUser,
		Password:    c.Config.DBPassword,
		Database:    c.Config.DBDatabase,
		SSLMode:     c.Config.DBSSLMode,
		Driver:      c.Config.DBDriver,
		Environment: c.Config.Environment,
	}
	c.Conn = NewConnection(&dbConfig)
	log.Println("✅ Database connected")
}
