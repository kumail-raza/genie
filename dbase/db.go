package dbase

import (
	"github.com/mohuk/genie/models"
)

// Store database store interface
type Store interface {
	GetDatabases() ([]models.Database, error)
	GetTables(string) ([]models.Table, error)
	GetColumns(db string, table string) ([]models.Column, error)
}

// NewStore creates a new store for the database
func NewStore(host string, port int, username, password string) Store {
	return &MSSqlDatabase{
		Host:     host,
		Port:     port,
		User:     username,
		Password: password,
	}
}
