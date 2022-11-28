package domain

const (
	SchemaTypeObject = "object"
)

// Schema is the structure that represents the schema section of an OpenAPI 3 contract
type Schema struct {
	Description string              `json:"description" yaml:"description"`
	Type        string              `json:"type" yaml:"type"`
	Required    []string            `json:"required" yaml:"required"`
	Properties  map[string]Property `json:"properties" yaml:"properties"`
}
