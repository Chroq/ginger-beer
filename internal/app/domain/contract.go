package domain

// Contract is the structure that represents an OpenAPI 3 contract
type Contract struct {
	OpenAPI    string `json:"openapi" yaml:"openAPI"`
	Info       Info   `json:"info" yaml:"info"`
	Servers    []Server
	Paths      map[string]Path
	Components Component
}
