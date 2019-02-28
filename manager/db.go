package manager

import "github.com/mohuk/genie/httpio"

// GetDatabases ...
func (g *genieManager) GetDatabases() ([]httpio.Database, error) {

	dbs, err := g.store.GetDatabases()
	if err != nil {
		return nil, err
	}
	databases := make([]httpio.Database, len(dbs))
	for j, db := range dbs {
		databases[j] = httpio.UnmarshallDB(db)
	}
	return databases, nil
}
