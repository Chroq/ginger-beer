package cli

import "go-openapi_builder/internal/app/adapter/service"

// Generator is the entry point for the cli application
type Generator struct {
	Config *service.Config
}

// NewGenerator creates a new instance of the Generator
func NewGenerator() *Generator {
	return &Generator{
		Config: service.NewConfig(),
	}
}

// Build executes the application
func (g *Generator) Build() error {

	return nil
}
