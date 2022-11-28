package factory

import (
	"ginger-beer/internal/app/adapter/service"
	"ginger-beer/internal/app/domain"
)

func BuildSchemaBySQLTable(table service.SQLTable) domain.Schema {
	return domain.Schema{
		Description: "",
		Type:        domain.SchemaTypeObject,
		Required:    nil,
		Properties:  BuildPropertiesBySQLTable(table),
	}
}

func BuildPropertiesBySQLTable(table service.SQLTable) map[string]domain.Property {
	properties := make(map[string]domain.Property, len(table.Fields))
	for i := range table.Fields {
		properties[table.Fields[i].Name] = domain.Property{
			Type:      table.Fields[i].Type,
			MaxLength: table.Fields[i].Size,
		}
	}
	return properties
}
