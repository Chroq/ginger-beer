package cli_test

import (
	"go-openapi_builder/internal/app/adapter/cli"
	"go-openapi_builder/internal/app/adapter/repository"
	"go-openapi_builder/internal/app/adapter/service"
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGeneratorNewGenerator(t *testing.T) {
	os.Args = []string{"main", "-c", "postgresql://localhost/postgres", "-o", "../../../../testdata/output", "-t", "clean", "-f", "yaml"}

	generator, err := cli.NewGenerator()
	td.Cmp(t, err, nil)

	td.Cmp(t, generator, &cli.Generator{
		SqlRepository: &repository.SqlRepository{
			DB: generator.SqlRepository.DB,
		},
		Config: &service.Config{
			Connection: "postgresql://localhost/postgres",
			Output:     "../../../../testdata/output",
			Type:       "clean",
			Driver:     "postgres",
			Format:     "yaml",
		},
	})

	if err := os.Remove("../../../../testdata/output"); err != nil {
		td.CmpNoError(t, err)
	}
}
