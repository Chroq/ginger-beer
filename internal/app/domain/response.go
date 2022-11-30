package domain

const (
	ContentTypeJSON = "application/json"
)

type Content struct {
	Schema Schema `json:"schema" yaml:"schema"`
}

type Response struct {
	Description string             `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]Content `json:"content,omitempty" yaml:"content,omitempty"`
}
