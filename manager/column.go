package manager

import (
	"fmt"

	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/formly"
)

// GetColumns ...
func (g *genieManager) GetColumns(dbname, tableName string) (*formly.TableForm, error) {

	columns, err := g.store.GetColumns(dbname, tableName)
	if err != nil {
		return nil, err
	}
	if len(columns) == 0 {
		return nil, errors.NewErrNoRows(fmt.Sprintf("table %s does not exist", tableName))
	}
	tf := formly.TableForm{
		TableName: tableName,
		Template:  make([]formly.Template, len(columns)),
	}
	for j, col := range columns {
		tf.Template[j] = g.formly.MapColumn(col)
	}
	return &tf, nil
}
