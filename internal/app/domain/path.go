package domain

const (
	OperationGet    = "get"
	OperationIndex  = "index"
	OperationPost   = "post"
	OperationPut    = "put"
	OperationDelete = "delete"
	URIPartID       = "/{id}"
)

func GetGlobalScopeVerbs(globalURI, unitaryURI string) map[string]string {
	return map[string]string{
		OperationIndex:  globalURI,
		OperationPost:   globalURI,
		OperationGet:    unitaryURI,
		OperationPut:    unitaryURI,
		OperationDelete: unitaryURI,
	}
}

// Path is the structure that represents the paths section of an OpenAPI 3 contract
type Path struct {
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string           `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Tags        []string         `json:"tags,omitempty" yaml:"tags,omitempty"`
	Parameters  []Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[int]Response `json:"responses" yaml:"responses"`
}
