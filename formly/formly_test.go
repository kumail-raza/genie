package formly

import (
	"database/sql"
	"testing"

	"github.com/mohuk/genie/models"
)

func TestHasMaxLength(t *testing.T) {
	cases := []struct {
		Expected bool
		Input    string
	}{
		{Expected: true, Input: "15"},
		{Expected: false, Input: "-15"},
	}

	for _, c := range cases {
		_, valid := hasValidMaxLength(c.Input)
		if valid != c.Expected {
			t.Fail()
		}
	}
}

func TestMapColumn(t *testing.T) {
	colName := "colName"
	length := "15"
	Nullable := "YES"
	someT := "text"
	c := models.Column{
		Name: sql.NullString{
			String: colName,
		},
		Length: sql.NullString{
			String: length,
		},
		Type: sql.NullString{
			String: someT,
		},
		Nullable: sql.NullString{
			String: Nullable,
		},
		Default: sql.NullString{},
	}
	m := NewFormlyMapper().MapColumn(c)
	if m.TemplateOps.Type != someT {
		t.Fail()
		return
	}
	if m.TemplateOps.Max != 15 {
		t.Fail()
		return
	}
	if m.TemplateOps.Required != true {
		t.Fail()
		return
	}
}
