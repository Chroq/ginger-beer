package usecase_test

import (
	"ginger-beer/internal/app/application/usecase"
	"ginger-beer/internal/app/domain"
	"ginger-beer/internal/app/domain/repository"
	"ginger-beer/internal/app/domain/repository/repository_mock"
	"ginger-beer/internal/app/domain/valueobject"
	"ginger-beer/testdata/tu"
	"testing"

	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/mock"
)

func TestContractUseCase_BuildContract(t *testing.T) {
	type fields struct {
		ContractRepository func(t *testing.T) repository.IContractRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.Contract
		err    error
	}{
		{
			name: "should return a contract",
			fields: fields{
				ContractRepository: func(t *testing.T) repository.IContractRepository {
					repositoryMock := repository_mock.NewIContractRepository(t)

					repositoryMock.
						On("GetEntities", mock.Anything).
						Return(map[string][]*valueobject.Field{
							"test": {
								{
									Name:      "test_var",
									Type:      "string",
									MaxLength: tu.Ptr(255),
								},
							},
						}, nil).
						Once()

					return repositoryMock
				},
			},
			want: &domain.Contract{
				OpenAPI: "3.0.3",
				Info:    domain.Info{},
				Paths: map[string]map[string]domain.Path{
					"/tests": {
						"get": {
							Description: "",
							OperationID: "getTests",
							Tags:        []string{"test"},
							Parameters: []domain.Parameter{
								{
									Reference: "#/components/parameters/",
								},
							},
							Responses: map[int]domain.Response{
								200: {
									Description: "The request has succeeded",
									Content: map[string]domain.Content{
										"application/json": {
											Schema: domain.Schema{
												Reference: "#/components/schemas/output.Test",
											},
										},
									},
								},
							},
						},
						"post": {
							Description: "",
							OperationID: "postTest",
							Tags:        []string{"test"},
							Parameters: []domain.Parameter{
								{
									Reference: "#/components/parameters/",
								},
							},
							Responses: map[int]domain.Response{
								201: {
									Description: "The request has been fulfilled and has resulted in one or more new resources being created",
									Content: map[string]domain.Content{
										"application/json": {
											Schema: domain.Schema{
												Reference: "#/components/schemas/output.Test",
											},
										},
									},
								},
							},
						},
					}, "/tests/{id}": {
						"get": {
							Description: "",
							OperationID: "getTest",
							Tags:        []string{"test"},
							Parameters: []domain.Parameter{
								{
									Reference: "#/components/parameters/",
								},
							},
							Responses: map[int]domain.Response{
								200: {
									Description: "The request has succeeded",
									Content: map[string]domain.Content{
										"application/json": {
											Schema: domain.Schema{
												Reference: "#/components/schemas/output.Test",
											},
										},
									},
								},
							},
						},
						"put": {
							Description: "",
							OperationID: "putTest",
							Tags:        []string{"test"},
							Parameters: []domain.Parameter{
								{
									Reference: "#/components/parameters/",
								},
							},
							Responses: map[int]domain.Response{
								200: {
									Description: "The request has succeeded",
									Content: map[string]domain.Content{
										"application/json": {
											Schema: domain.Schema{
												Reference: "#/components/schemas/output.Test",
											},
										},
									},
								},
							},
						},
						"delete": {
							Description: "",
							OperationID: "deleteTest",
							Tags:        []string{"test"},
							Parameters: []domain.Parameter{
								{
									Reference: "#/components/parameters/",
								},
							},
							Responses: map[int]domain.Response{
								204: {
									Description: "The server has successfully fulfilled the request and that there is no content to send in the response payload body",
									Content:     nil,
								},
							},
						},
					},
				},
				Servers: []domain.Server{
					{
						URL:         "http://localhost:8080",
						Description: "Local development server",
					},
				},
				Component: domain.Component{
					Schema: map[string]domain.Schema{
						"input.Test": {
							Type: "object",
							Properties: map[string]domain.Property{
								"test_var": {
									MaxLength: tu.Ptr(255),
									Type:      "string",
								},
							},
						},
						"output.Test": {
							Type: "object",
							Properties: map[string]domain.Property{
								"test_var": {
									MaxLength: tu.Ptr(255),
									Type:      "string",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase.ContractUseCase{
				ContractRepository: tt.fields.ContractRepository(t),
			}
			got, err := u.BuildContract()
			td.Cmp(t, err, tt.err)
			td.Cmp(t, got, tt.want)
		})
	}
}
