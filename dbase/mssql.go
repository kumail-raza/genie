package dbase

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/models"
)

// MSSqlDatabase class as a struct
type MSSqlDatabase struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	Conn         *sql.DB
}

// GetDatabases ..
func (db *MSSqlDatabase) GetDatabases() ([]models.Database, error) {
	connStr := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s", db.Host, db.Port, db.User, db.Password)
	conn, err := sql.Open("mssql", connStr)
	if err != nil {
		return nil, errors.NewErrDbConn(err.Error())
	}
	defer conn.Close()
	q := fmt.Sprint("SELECT name FROM master.dbo.sysdatabases WHERE name NOT IN ('master', 'tempdb', 'model', 'msdb');")
	rows, err := conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var databases []models.Database
	for rows.Next() {
		var db models.Database
		err := rows.Scan(&db.DBName)
		if err != nil {
			return nil, err
		}
		databases = append(databases, db)
	}
	return databases, nil
}

// Connect to initialize database connection
func (db *MSSqlDatabase) Connect(dbName string) error {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", db.Host, db.User, db.Password, db.Port, dbName)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return errors.NewErrDbConn(err.Error())
	}
	db.Conn = conn
	return nil
}

// Close connection to DB
func (db *MSSqlDatabase) Close() error {
	return db.Conn.Close()
}

// GetTables return the tables for database
func (db *MSSqlDatabase) GetTables(dbName string) ([]models.Table, error) {
	err := db.Connect(dbName)
	if err != nil {
		return nil, err
	}
	qry := fmt.Sprintf("SELECT DISTINCT TABLE_NAME, TABLE_CATALOG, TABLE_SCHEMA, TABLE_TYPE FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_CATALOG = '%s' AND TABLE_TYPE = 'BASE TABLE'", dbName)
	rows, err := db.Conn.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables = []models.Table{}
	for rows.Next() {
		var tbl = models.Table{}
		err := rows.Scan(&tbl.Name, &tbl.Catalog, &tbl.Schema, &tbl.Type)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tbl)
	}
	return tables, nil
}

// GetColumns returns the columns for a table
func (db *MSSqlDatabase) GetColumns(dbname, table string) ([]models.Column, error) {
	err := db.Connect(dbname)
	if err != nil {
		return nil, err
	}
	qry := fmt.Sprintf("SELECT DISTINCT COLUMN_NAME, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH FROM INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = '%s' AND TABLE_CATALOG = '%s'", table, dbname)
	rows, err := db.Conn.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns []models.Column
	for rows.Next() {
		var col models.Column
		err := rows.Scan(&col.Name, &col.Default, &col.Nullable, &col.Type, &col.Length)
		if err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}
	return columns, nil
}
