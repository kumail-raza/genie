package formly

import "github.com/mohuk/genie/models"

type formly struct{}

// Mapper ..
type Mapper interface {
	MapColumn(c models.Column) Template
}

func NewFormlyMapper() Mapper {
	return &formly{}
}
