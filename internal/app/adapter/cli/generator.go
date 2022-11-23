package cli

import "go-openapi_builder/internal/app/adapter/service"

// Generator is the entry point for the cli application
type Generator struct {
	Config *service.Config
}

// NewGenerator creates a new instance of the Generator
func NewGenerator() (*Generator, error) {
	config, err := service.NewConfig()
	if err != nil {
		return nil, err
	}

	return &Generator{
		Config: config,
	}, nil
}

// Build executes the application
func (g *Generator) Build() error {

	return nil
}
