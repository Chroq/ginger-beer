package factory

import (
	"ginger-beer/internal/app/domain"

	"github.com/tangzero/inflector"
)

func BuildPathsByEntities(entities, verbs []string) map[string]map[string]domain.Path {
	paths := make(map[string]map[string]domain.Path, len(entities))
	for i := range entities {
		newURI := "/" + inflector.Pluralize(inflector.Dasherize(entities[i]))
		paths[newURI] = make(map[string]domain.Path, len(verbs))
		for j := range verbs {
			paths[newURI][verbs[j]] = domain.Path{
				OperationID: verbs[j] + inflector.Camelize(entities[i]),
				Tags: []string{
					entities[i],
				},
				Parameters: []domain.Parameter{
					{
						Reference: "#/components/parameters/",
					},
				},
			}
		}
	}

	return paths
}
