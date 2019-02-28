package manager

import "github.com/mohuk/genie/httpio"

// GetTables ...
func (g *genieManager) GetTables(dbname string) ([]httpio.Table, error) {

	tbs, err := g.store.GetTables(dbname)
	if err != nil {
		return nil, err
	}
	tables := make([]httpio.Table, len(tbs))
	for j, t := range tbs {
		tables[j] = httpio.UnmarshallTable(t)
	}
	return tables, nil
}
