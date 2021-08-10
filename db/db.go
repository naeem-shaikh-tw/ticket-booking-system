package db

import (
	"database/sql"
	"fmt"
)

var DbPool *sql.DB

type Db struct {
	Name     string
	Host     string
	Port     int
	UserName string
	Password string
}

func (dbConfig Db) Connect() *sql.DB {
	pool, error := sql.Open("postgres", dbConfig.ConnectionUrl())
	if error != nil {
		panic(fmt.Sprint("Error connecting postgres", error))
	}
	return pool
}

func (dbConfig Db) ConnectionUrl() string {
	connectionUrl := "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	return fmt.Sprintf(connectionUrl, dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}
