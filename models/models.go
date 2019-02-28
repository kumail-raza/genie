package models

import "database/sql"

// Column is a column
type Column struct {
	Name     sql.NullString
	Type     sql.NullString
	Nullable sql.NullString
	Default  sql.NullString
	Length   sql.NullString
}

// Table is a table
type Table struct {
	Catalog sql.NullString
	Schema  sql.NullString
	Name    sql.NullString
	Type    sql.NullString
}

type Database struct {
	DBName sql.NullString `json:"dbName,omitempty"`
}
