package formly

import (
	"fmt"
	"strconv"

	"github.com/mohuk/genie/models"
)

func (f *formly) MapColumn(col models.Column) Template {

	t := Template{
		Key:  col.Name.String,
		Type: "input",
		TemplateOps: TemplateOpts{
			Type:        Type(col.Type.String),
			PlaceHolder: fmt.Sprintf("Enter %s...", col.Name.String),
			Required:    Required(col.Nullable.String),
		},
	}
	integer, valid := hasValidMaxLength(col.Length.String)
	if !valid {
		return t
	}
	t.TemplateOps.Max = *integer
	return t

}

func hasValidMaxLength(s string) (*int, bool) {
	if s == "" {
		return nil, false
	}
	intified, err := strconv.Atoi(s)
	if err != nil {
		return nil, false
	}
	if intified <= 0 {
		return nil, false
	}
	return &intified, true
}
