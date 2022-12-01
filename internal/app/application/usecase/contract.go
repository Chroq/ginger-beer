package usecase

import (
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/factory"
	"ginger-beer/internal/app/domain/repository"
	"ginger-beer/internal/app/domain/valueobject"
	"net/http"

	"github.com/tangzero/inflector"
)

type ContractUseCase struct {
	ContractRepository repository.IContractRepository
}

func (u *ContractUseCase) BuildContract() (*domain.Contract, error) {
	entities, err := u.ContractRepository.GetEntities()
	if err != nil {
		return nil, err
	}

	schemas := make(map[string]domain.Schema, len(entities))
	for entity := range entities {
		outputEntity := domain.ReferencePrefixOutput + inflector.Camelize(entity)
		schemas[outputEntity] = factory.BuildSchemaByEntity(entities[entity])
		inputEntity := domain.ReferencePrefixInput + inflector.Camelize(entity)
		schemas[inputEntity] = factory.BuildSchemaByEntity(entities[entity])
	}

	return &domain.Contract{
		Servers: []domain.Server{
			domain.GetDefaultServer(),
		},
		OpenAPI: domain.DefaultOpenAPIVersion,
		Component: domain.Component{
			Schema: schemas,
		},
		Paths: BuildPathsByEntities(entities, []string{
			domain.OperationGet,
			domain.OperationIndex,
			domain.OperationPost,
			domain.OperationPut,
			domain.OperationDelete,
		}),
	}, nil
}

func BuildPathsByEntities(entities map[string][]*valueobject.Field, verbs []string) map[string]map[string]domain.Path {
	paths := make(map[string]map[string]domain.Path, len(entities))
	for entity := range entities {
		globalURI := getBaseURI(entity)
		unitaryURI := globalURI + domain.URIPartID

		globalVerbs := domain.GetGlobalScopeVerbs(globalURI, unitaryURI)
		const NumberOfGlobalVerbs, NumberOfUnitaryVerbs = 2, 3

		paths[globalURI] = make(map[string]domain.Path, NumberOfGlobalVerbs)
		paths[unitaryURI] = make(map[string]domain.Path, NumberOfUnitaryVerbs)
		for j := range verbs {
			outputSchemaReference := domain.GetOutputSchemaReference(entity)
			verb := verbs[j]
			if verb == domain.OperationIndex {
				verb = domain.OperationGet
			}

			paths[globalVerbs[verb]][verb] = domain.Path{
				OperationID: verb + inflector.Camelize(entity),
				Tags: []string{
					entity,
				},
				Parameters: []domain.Parameter{
					{
						Reference: domain.BaseReferenceComponentsParameter,
					},
				},
				Responses: map[int]domain.Response{
					getStatusByVerb(verb): getResponseByVerb(verb, outputSchemaReference),
				},
			}
		}
	}

	return paths
}

func getBaseURI(entity string) string {
	newURI := "/" + inflector.Pluralize(inflector.Dasherize(entity))
	return newURI
}

func getResponseByVerb(verb, outputSchemaReference string) domain.Response {
	var response domain.Response
	switch verb {
	case domain.OperationDelete:
		response = domain.Response{
			Description: domain.Default204Description,
		}
	case domain.OperationIndex:
		response = domain.Response{
			Description: domain.Default200Description,
			Content: map[string]domain.Content{
				domain.ContentTypeJSON: {
					Schema: domain.Schema{
						Reference: outputSchemaReference,
					},
				},
			},
		}
	default:
		response = domain.Response{
			Description: domain.Default200Description,
			Content: map[string]domain.Content{
				domain.ContentTypeJSON: {
					Schema: domain.Schema{
						Reference: outputSchemaReference,
					},
				},
			},
		}
	}

	return response
}

func getStatusByVerb(verb string) int {
	var status int
	switch verb {
	case domain.OperationDelete:
		status = http.StatusNoContent
	case domain.OperationPost:
		status = http.StatusCreated
	default:
		status = http.StatusOK
	}
	return status
}
