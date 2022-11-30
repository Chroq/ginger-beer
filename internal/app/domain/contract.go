package domain

const (
	DefaultOpenAPIVersion = "3.0.3"
)

// Contract is the structure that represents an OpenAPI 3 contract
type Contract struct {
	OpenAPI   string                     `json:"openapi" yaml:"openapi"`
	Info      Info                       `json:"info" yaml:"info"`
	Servers   []Server                   `json:"servers" yaml:"servers"`
	Paths     map[string]map[string]Path `json:"paths" yaml:"paths"`
	Component Component                  `json:"components" yaml:"components"`
}
