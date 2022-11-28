package main

import (
	"fmt"
	"ginger-beer/internal/app/adapter/cli"
	"ginger-beer/internal/app/application/usecase"
)

func main() {
	if generator, err := cli.NewGenerator(); err == nil {
		contractUseCase := usecase.ContractUseCase{
			ComponentRepository: generator.SQLRepository,
		}
		if contract, err := contractUseCase.BuildContract(); err == nil {
			fmt.Println(contract)
		} else {
			fmt.Printf("error: %s", err)
		}
	} else {
		fmt.Printf("error: %s", err)
	}
}
