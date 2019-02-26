package dbase

import (
	"database/sql"

	"github.com/mohuk/genie/models"
)

type mockstore struct{}

func NewMockStore() Store {
	return &mockstore{}
}

func (m *mockstore) GetDatabases() ([]models.Database, error) {
	return []models.Database{
		models.Database{DBName: sql.NullString{
			String: "db",
		}},
	}, nil
}
func (m *mockstore) GetTables(string) ([]models.Table, error) {
	return []models.Table{
		models.Table{
			Name: sql.NullString{
				String: "table",
			},
		},
	}, nil
}
func (m *mockstore) GetColumns(db string, table string) ([]models.Column, error) {
	return []models.Column{
		models.Column{Name: sql.NullString{
			String: "column",
		}},
	}, nil
}
