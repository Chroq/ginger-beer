package domain

const (
	OperationGet    = "get"
	OperationPost   = "post"
	OperationPut    = "put"
	OperationDelete = "delete"
)

// Path is the structure that represents the paths section of an OpenAPI 3 contract
type Path struct {
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string      `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Tags        []string    `json:"tags,omitempty" yaml:"tags,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
