package cli

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"ginger-beer/internal/app/adapter/repository"
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/application/usecase"
	"gopkg.in/yaml.v3"

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

// Generate creates a new contract
func (g *Generator) Generate() error {
	contractUseCase := usecase.ContractUseCase{
		ContractRepository: g.SQLRepository,
	}
	if contract, err := contractUseCase.BuildContract(); err == nil {
		if g.Config.Format == service.PermittedFormatJSON {
			if marshal, err := json.Marshal(contract); err == nil {
				fmt.Println(string(marshal))
			} else {
				return err
			}
		} else if g.Config.Format == service.PermittedFormatYAML {
			if marshal, err := yaml.Marshal(contract); err == nil {
				fmt.Println(string(marshal))
			} else {
				return err
			}
		}
	}

	return nil
}
