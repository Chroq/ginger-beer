package cli

import (
	"database/sql"
	"go-openapi_builder/internal/app/adapter/repository"
	"go-openapi_builder/internal/app/adapter/service"

	_ "github.com/lib/pq"
)

// Generator is the entry point for the cli application
type Generator struct {
	Config        *service.Config
	SqlRepository *repository.SqlRepository
}

// NewGenerator creates a new instance of the Generator
func NewGenerator() (*Generator, error) {
	config, err := service.NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(config.Driver, config.Connection)
	if err != nil {
		return nil, err
	}

	return &Generator{
		Config: config,
		SqlRepository: &repository.SqlRepository{
			DB: db,
		},
	}, nil
}

// Build executes the application
func (g *Generator) Build() error {
	_, err := g.SqlRepository.GetTableNames()
	if err != nil {
		return err
	}

	return nil
}
