package manager

import (
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/formly"
	"github.com/mohuk/genie/httpio"
)

type GenieManager interface {
	GetColumns(dbname, tableName string) (*formly.TableForm, error)
	GetDatabases() ([]httpio.Database, error)
	GetTables(dbname string) ([]httpio.Table, error)
}

type genieManager struct {
	store  dbase.Store
	formly formly.Mapper
}

func NewGenieManager(s dbase.Store, m formly.Mapper) GenieManager {
	return &genieManager{store: s, formly: m}
}
