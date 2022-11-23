package domain

// Path is the structure that represents the paths section of an OpenAPI 3 contract
type Path struct {
	Ref string `json:"$ref" yaml:"$ref"`
}
