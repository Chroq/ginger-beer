package domain

const (
	ContentTypeJSON       = "application/json"
	Default200Description = "The request has succeeded"
	Default204Description = "The server has successfully fulfilled the request and that there is no content to send in the response payload body"
	Defautl201Description = "The request has been fulfilled and has resulted in one or more new resources being created"
)

type Content struct {
	Schema Schema `json:"schema" yaml:"schema"`
}

type Response struct {
	Description string             `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]Content `json:"content,omitempty" yaml:"content,omitempty"`
}
