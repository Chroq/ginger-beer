package domain

// Server is the structure that represents the servers section of an OpenAPI 3 contract
type Server struct {
	URL         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
}
