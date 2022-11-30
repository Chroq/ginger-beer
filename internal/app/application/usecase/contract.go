package usecase

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/factory"
	"ginger-beer/internal/app/domain/repository"
)

type ContractUseCase struct {
	ComponentRepository repository.IComponentRepository
}

func (u *ContractUseCase) BuildContract() (*domain.Contract, error) {
	component, err := u.ComponentRepository.GetComponent()
	if err != nil {
		return nil, err
	}

	entities, err := u.ComponentRepository.GetEntities()
	if err != nil {
		return nil, err
	}

	return &domain.Contract{
		OpenAPI:   domain.DefaultOpenAPIVersion,
		Component: *component,
		Paths: factory.BuildPathsByEntities(entities, []string{
			domain.OperationGet,
			domain.OperationPost,
			domain.OperationPut,
			domain.OperationDelete,
		}),
	}, nil
}
