package cli_test

import (
	"go-openapi_builder/internal/app/adapter/cli"
	"go-openapi_builder/internal/app/adapter/service"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGeneratorBuild(t *testing.T) {
	generator := cli.Generator{}

	td.CmpNoError(t, generator.Build())
}

func TestGeneratorNewGenerator(t *testing.T) {
	generator, err := cli.NewGenerator()
	td.Cmp(t, err, nil)
	td.Cmp(t, generator, &cli.Generator{
		Config: &service.Config{
			File:   "model.go",
			Output: ".",
			Type:   "basic",
			Format: "json",
		},
	})
}
