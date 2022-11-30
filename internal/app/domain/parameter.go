package domain

const (
	BaseReferenceComponentsParameter = "#/components/parameters/"
)

type Parameter struct {
	Reference string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
