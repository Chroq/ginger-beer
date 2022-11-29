package factory

import (
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/adapter/translator"
	"ginger-beer/internal/app/domain"
)

func BuildSchemaBySQLTable(table service.SQLTable) (*domain.Schema, error) {
	properties, err := BuildPropertiesBySQLTable(table)
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

func BuildPropertiesBySQLTable(table service.SQLTable) (map[string]domain.Property, error) {
	properties := make(map[string]domain.Property, len(table.Fields))
	for i := range table.Fields {
		openAPIType, err := translator.SQLToOpenAPIScalarTypes(table.Fields[i].Type)
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
