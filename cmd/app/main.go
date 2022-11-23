package main

import (
	"fmt"
	"go-openapi_builder/internal/app/adapter/cli"
)

func main() {
	generator, err := cli.NewGenerator()
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	if err := generator.Build(); err != nil {
		fmt.Printf("error: %s", err)
	}
}
