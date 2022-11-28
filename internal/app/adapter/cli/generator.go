package cli

import (
	"database/sql"
	"ginger-beer/internal/app/adapter/repository"
	"ginger-beer/internal/app/adapter/service"

	_ "github.com/lib/pq"
)

// Generator is the entry point for the cli application
type Generator struct {
	Config        *service.Config
	SQLRepository *repository.SQLRepository
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
		SQLRepository: &repository.SQLRepository{
			DB: db,
		},
	}, nil
}
