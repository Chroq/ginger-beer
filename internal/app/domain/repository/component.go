package repository

import "ginger-beer/internal/app/domain"

type IComponentRepository interface {
	GetComponent() (*domain.Component, error)
	GetEntities() ([]string, error)
}
