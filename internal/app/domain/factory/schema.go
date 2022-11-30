package factory

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/valueobject"
)

func BuildSchemaByEntity(fields []*valueobject.Field) domain.Schema {
	return domain.Schema{
		Type:       domain.SchemaTypeObject,
		Properties: BuildPropertiesByEntity(fields),
	}
}

func BuildPropertiesByEntity(fields []*valueobject.Field) map[string]domain.Property {
	properties := make(map[string]domain.Property, len(fields))
	for i := range fields {
		properties[fields[i].Name] = domain.Property{
			Type:      fields[i].Type,
			MaxLength: fields[i].MaxLength,
		}
	}
	return properties
}
