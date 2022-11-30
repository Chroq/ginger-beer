package domain

const (
	BaseReferenceComponentsSchemas = "#/components/schemas/"
	SchemaTypeObject               = "object"
)

// Schema is the structure that represents the schema section of an OpenAPI 3 contract
type Schema struct {
	Reference   string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string              `json:"type,omitempty" yaml:"type,omitempty"`
	Required    []string            `json:"required,omitempty" yaml:"required,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
}
