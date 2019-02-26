package manager

import (
	"fmt"

	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/models"
)

type GenieManager interface {
	GetColumns(dbname, tableName string) (*models.TableForm, error)
	GetDatabases() ([]httpio.Database, error)
	GetTables(dbname string) ([]httpio.Table, error)
}

type genieManager struct {
	store dbase.Store
}

func NewGenieManager(s dbase.Store) GenieManager {
	return &genieManager{store: s}
}

// GetColumns ...
func (g *genieManager) GetColumns(dbname, tableName string) (*models.TableForm, error) {

	columns, err := g.store.GetColumns(dbname, tableName)
	if err != nil {
		return nil, err
	}
	tf := models.TableForm{TableName: tableName, Template: []models.Template{}}
	for _, col := range columns {
		tf.Template = append(tf.Template, models.Template{
			Key:  col.Name.String,
			Type: "input",
			TemplateOps: models.TemplateOpts{
				Type:        col.Type.String,
				PlaceHolder: fmt.Sprintf("Enter %s...", col.Name.String),
			},
		})
	}
	return &tf, nil
}

// GetDatabases ...
func (g *genieManager) GetDatabases() ([]httpio.Database, error) {

	dbs, err := g.store.GetDatabases()
	if err != nil {
		return nil, err
	}
	var resp []httpio.Database
	for _, x := range dbs {
		resp = append(resp, httpio.UnmarshallDB(x))
	}
	return resp, nil
}

// GetTables ...
func (g *genieManager) GetTables(dbname string) ([]httpio.Table, error) {

	tables, err := g.store.GetTables(dbname)
	if err != nil {
		return nil, err
	}
	var resp []httpio.Table
	for _, t := range tables {
		resp = append(resp, httpio.UnmarshallTable(t))
	}
	return resp, nil
}
