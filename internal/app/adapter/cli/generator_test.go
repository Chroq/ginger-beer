package cli_test

import (
	"go-openapi_builder/internal/app/adapter/cli"
	"go-openapi_builder/internal/app/adapter/service"
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGeneratorBuild(t *testing.T) {
	generator := cli.Generator{}

	td.CmpNoError(t, generator.Build())
}

func TestGeneratorNewGenerator(t *testing.T) {
	os.Args = []string{"main", "-c", "postgresql://localhost/postgres", "-o", "../../../../testdata/output", "-t", "clean", "-f", "yaml"}

	generator, err := cli.NewGenerator()
	td.Cmp(t, err, nil)
	td.Cmp(t, generator, &cli.Generator{
		Config: &service.Config{
			Connection: "postgresql://localhost/postgres",
			Output:     "../../../../testdata/output",
			Type:       "clean",
			Format:     "yaml",
		},
	})

	if err := os.Remove("../../../../testdata/output"); err != nil {
		td.CmpNoError(t, err)
	}
}