package main

import (
	"fmt"
	"go-openapi_builder/internal/app/adapter"
)

func main() {
	err := adapter.NewGenerator().Build()
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
