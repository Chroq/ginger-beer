package factory

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/valueobject"
	"net/http"

	"github.com/tangzero/inflector"
)

func BuildPathsByEntities(entities map[string][]*valueobject.Field, verbs []string) map[string]map[string]domain.Path {
	paths := make(map[string]map[string]domain.Path, len(entities))
	for entity := range entities {
		newURI := "/" + inflector.Pluralize(inflector.Dasherize(entity))
		paths[newURI] = make(map[string]domain.Path, len(verbs))
		for j := range verbs {
			outputSchemaReference := domain.GetOutputSchemaReference(entity)
			paths[newURI][verbs[j]] = domain.Path{
				OperationID: verbs[j] + inflector.Camelize(entity),
				Tags: []string{
					entity,
				},
				Parameters: []domain.Parameter{
					{
						Reference: domain.BaseReferenceComponentsParameter,
					},
				},
				Responses: map[int]domain.Response{
					http.StatusOK: {
						Description: domain.Default200Description,
						Content: map[string]domain.Content{
							domain.ContentTypeJSON: {
								Schema: domain.Schema{
									Reference: outputSchemaReference,
								},
							},
						},
					},
				},
			}
		}
	}

	return paths
}
