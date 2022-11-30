package factory

import (
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/adapter/translator"
	"ginger-beer/internal/app/domain/valueobject"
)

func PgFieldsToFields(pgFields []service.PgField) ([]*valueobject.Field, error) {
	fields := make([]*valueobject.Field, len(pgFields))
	for i := range pgFields {
		field, err := PgFieldToField(pgFields[i])
		if err != nil {
			return nil, err
		}
		fields[i] = field
	}
	return fields, nil
}

func PgFieldToField(pgField service.PgField) (*valueobject.Field, error) {
	types, err := translator.PgSQLToOpenAPITypes(pgField.Type)
	if err != nil {
		return nil, err
	}
	return &valueobject.Field{
		Name:      pgField.Name,
		Type:      types,
		MaxLength: pgField.Size,
	}, nil
}
