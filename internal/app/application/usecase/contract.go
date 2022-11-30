package usecase

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/factory"
	"ginger-beer/internal/app/domain/repository"
)

type ContractUseCase struct {
	ContractRepository repository.IContractRepository
}

func (u *ContractUseCase) BuildContract() (*domain.Contract, error) {
	entities, err := u.ContractRepository.GetEntities()
	if err != nil {
		return nil, err
	}

	schemas := make(map[string]domain.Schema, len(entities))
	for entity := range entities {
		schemas[entity] = factory.BuildSchemaByEntity(entities[entity])
	}

	return &domain.Contract{
		Servers: []domain.Server{
			domain.GetDefaultServer(),
		},
		OpenAPI: domain.DefaultOpenAPIVersion,
		Component: domain.Component{
			Schema: schemas,
		},
		Paths: factory.BuildPathsByEntities(entities, []string{
			domain.OperationGet,
			domain.OperationPost,
			domain.OperationPut,
			domain.OperationDelete,
		}),
	}, nil
}
