package formly

import "github.com/mohuk/genie/models"

type mockformly struct{}

func (m *mockformly) MapColumn(c models.Column) Template {
	return Template{}
}

func NewMockFormly() Mapper {
	return &mockformly{}
}
