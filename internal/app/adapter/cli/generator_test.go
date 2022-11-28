package cli_test

import (
	"ginger-beer/internal/app/adapter/cli"
	"ginger-beer/internal/app/adapter/repository"
	"ginger-beer/internal/app/adapter/service"
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGeneratorNewGenerator(t *testing.T) {
	os.Args = []string{"main", "-c", "postgresql://localhost/postgres", "-o", "../../../../testdata/output", "-t", "clean", "-f", "yaml"}

	generator, err := cli.NewGenerator()
	td.Cmp(t, err, nil)

	td.Cmp(t, generator, &cli.Generator{
		SQLRepository: &repository.SQLRepository{
			DB: generator.SQLRepository.DB,
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
