package domain

// Component is the structure that represents the components section of an OpenAPI 3 contract
type Component struct {
	Schema map[string]Schema `json:"schemas" yaml:"schemas"`
}
