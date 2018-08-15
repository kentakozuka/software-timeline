package main

import (
	"fmt"
	"os"

	"github.com/kentakozuka/software-timeline/pkg/parser/cmd"
)

func main() {
	command := cmd.NewDefaultParserCommand()
	// pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	// pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	// logs.InitLogs()
	// defer logs.FlushLogs()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
