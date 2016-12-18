package main

import (
	"fmt"
	"os"

	"github.com/minamijoyo/api-cli-go-example/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
