package manager

import (
	"testing"

	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/formly"
)

func TestGetDatabases(t *testing.T) {

	s := dbase.NewMockStore()
	mapper := formly.NewMockFormly()
	m := NewGenieManager(s, mapper)
	dbs, err := m.GetDatabases()
	if err != nil {
		t.Fail()
	}
	if dbs[0].Name != "db" {
		t.Fail()
	}
}

func TestGetColumns(t *testing.T) {
	s := dbase.NewMockStore()
	mapper := formly.NewMockFormly()
	m := NewGenieManager(s, mapper)
	tName := "tableName"
	form, err := m.GetColumns("dbname", tName)
	if err != nil {
		t.Fail()
	}
	if form.TableName != tName {
		t.Fail()
	}
}

func TestGetTables(t *testing.T) {

	s := dbase.NewMockStore()
	mapper := formly.NewMockFormly()
	m := NewGenieManager(s, mapper)
	dbName := "dbname"
	tables, err := m.GetTables(dbName)
	if err != nil {
		t.Fail()
	}
	if tables[0].Name != "table" {
		t.Fail()
	}
}
