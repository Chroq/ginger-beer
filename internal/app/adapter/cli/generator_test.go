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
	td.Cmp(t, cli.NewGenerator(), &cli.Generator{
		Config: &service.Config{
			File:   "model.go",
			Output: ".",
			Type:   "basic",
			Format: "json",
		},
	})
}
