package main

import (
	"fmt"
	"os"

	"github.com/kentakozuka/software-timeline/pkg/gateway/cmd"
)

func main() {
	command := cmd.NewDefaultGatewayCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
