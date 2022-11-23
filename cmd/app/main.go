package main

import (
	"fmt"
	"go-openapi_builder/internal/app/adapter/cli"
)

func main() {
	err := cli.NewGenerator().Build()
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
