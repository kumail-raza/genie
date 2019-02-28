package formly

type TableForm struct {
	TableName string     `json:"tableName,omitempty"`
	Template  []Template `json:"formlyTemplate,omitempty"`
}
type Template struct {
	Key         string       `json:"key,omitempty"`
	Type        string       `json:"type,omitempty"`
	TemplateOps TemplateOpts `json:"templateOptions,omitempty"`
}
type TemplateOpts struct {
	Type        string `json:"type,omitempty"`
	Label       string `json:"label,omitempty"`
	PlaceHolder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Min         int    `json:"min,omitempty"`
	Max         int    `json:"max,omitempty"`
}

var typeMaps = map[string]string{
	"nvarchar":  "text",
	"datetime2": "date",
	"bigint":    "number",
	"geography": "number",
	"decimal":   "number",
	"date":      "date",
	"bit":       "checkbox",
	"varbinary": "checkbox",
	"int":       "number",
}

var requiredMap = map[string]bool{
	"NO":  false,
	"YES": true,
}

// Required ...
func Required(s string) bool {
	if v, ok := requiredMap[s]; ok {
		return v
	}
	return false
}

// Type ...
func Type(s string) string {
	if v, ok := typeMaps[s]; ok {
		return v
	}
	return "text"
}
