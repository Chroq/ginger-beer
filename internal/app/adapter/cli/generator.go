package cli

import (
	"database/sql"
	"fmt"
	"go-openapi_builder/internal/app/adapter/service"

	_ "github.com/lib/pq"
)

// Generator is the entry point for the cli application
type Generator struct {
	Config *service.Config
	DB     *sql.DB
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}(db)

	return &Generator{
		Config: config,
		DB:     db,
	}, nil
}

// Build executes the application
func (g *Generator) Build() error {

	return nil
}
