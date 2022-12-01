package repository

import (
	"ginger-beer/internal/app/domain/valueobject"
)

type IContractRepository interface {
	GetEntities() (map[string][]*valueobject.Field, error)
}
