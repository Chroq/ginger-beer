package main

import (
	"fmt"
	"ginger-beer/internal/app/adapter/cli"
)

func main() {
	if generator, err := cli.NewGenerator(); err == nil {
		if err := generator.Generate(); err != nil {
			fmt.Printf("error: %s", err)
		}
	} else {
		fmt.Printf("error: %s", err)
	}
}
