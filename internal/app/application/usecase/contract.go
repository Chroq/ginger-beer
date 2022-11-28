package usecase

import (
	"ginger-beer/internal/app/domain"
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

	return &domain.Contract{
		Component: *component,
	}, nil
}
