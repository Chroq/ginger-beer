package factory

import (
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/adapter/translator"
	"ginger-beer/internal/app/domain"
)

func BuildSchemaByPgTable(table service.PgTable) (*domain.Schema, error) {
	properties, err := BuildPropertiesByPgTable(table)
	if err != nil {
		return nil, err
	}
	return &domain.Schema{
		Description: "",
		Type:        domain.SchemaTypeObject,
		Required:    nil,
		Properties:  properties,
	}, nil
}

func BuildPropertiesByPgTable(table service.PgTable) (map[string]domain.Property, error) {
	properties := make(map[string]domain.Property, len(table.Fields))
	for i := range table.Fields {
		openAPIType, err := translator.PgSQLToOpenAPITypes(table.Fields[i].Type)
		if err != nil {
			return nil, err
		}
		properties[table.Fields[i].Name] = domain.Property{
			Type:      openAPIType,
			MaxLength: table.Fields[i].Size,
		}
	}
	return properties, nil
}
