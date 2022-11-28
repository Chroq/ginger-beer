package domain

type Property struct {
	MaxLength *int   `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Type      string `json:"type" yaml:"type"`
}
