package domain

// Contract is the structure that represents an OpenAPI 3 contract
type Contract struct {
	OpenAPI   string          `json:"openapi" yaml:"openAPI"`
	Info      Info            `json:"info" yaml:"info"`
	Component Component       `json:"components" yaml:"components"`
	Servers   []Server        `json:"servers" yaml:"servers"`
	Paths     map[string]Path `json:"paths" yaml:"paths"`
}
