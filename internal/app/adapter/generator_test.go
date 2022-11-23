package adapter_test

import (
	"go-openapi_builder/internal/app/adapter"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestGeneratorBuild(t *testing.T) {
	generator := adapter.Generator{}

	td.CmpNoError(t, generator.Build())
}

func TestGeneratorNewGenerator(t *testing.T) {
	td.Cmp(t, adapter.NewGenerator(), &adapter.Generator{})
}
