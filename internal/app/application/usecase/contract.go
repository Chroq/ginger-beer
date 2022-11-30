package usecase

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/factory"
	"ginger-beer/internal/app/domain/repository"

	"github.com/tangzero/inflector"
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
		outputEntity := domain.ReferencePrefixOutput + inflector.Camelize(entity)
		schemas[outputEntity] = factory.BuildSchemaByEntity(entities[entity])
		inputEntity := domain.ReferencePrefixInput + inflector.Camelize(entity)
		schemas[inputEntity] = factory.BuildSchemaByEntity(entities[entity])
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
