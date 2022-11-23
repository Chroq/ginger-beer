package domain

// Schema is the structure that represents the schema section of an OpenAPI 3 contract
type Schema struct {
	Ref string `json:"$ref" yaml:"$ref"`
}
