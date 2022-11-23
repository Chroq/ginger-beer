package domain

// Info is the object providing metadata about the API.
type Info struct {
	Title          string `json:"title" yaml:"title"`
	Description    string `json:"description" yaml:"description"`
	TermsOfService string `json:"termsOfService" yaml:"termsOfService"`
	Contact        struct {
		Name  string `json:"name" yaml:"name"`
		URL   string `json:"url" yaml:"url"`
		Email string `json:"email" yaml:"email"`
	} `json:"contact" yaml:"contact"`
	License struct {
		Name string `json:"name" yaml:"name"`
		URL  string `json:"url" yaml:"URL"`
	} `json:"license" yaml:"license"`
	Version string `json:"version" yaml:"version"`
}
