package cmd

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
)

func NewDefaultGatewayCommand() *cobra.Command {
	return NewDefaultGatewayCommandWithArgs(&defaultPluginHandler{}, os.Args, os.Stdin, os.Stdout, os.Stderr)
}

func NewDefaultGatewayCommandWithArgs(pluginHandler PluginHandler, args []string, in io.Reader, out, errout io.Writer) *cobra.Command {
	cmd := NewGatewayCommand(in, out, errout)

	if pluginHandler == nil {
		return cmd
	}
	return cmd
}

func NewDefaultParserGatewayWithArgs(ph PluginHandler, args []string, in io.Reader, out, errout io.Writer) *cobra.Command {
	cmd := NewGatewayCommand(in, out, errout)

	if ph == nil {
		return cmd
	}
	return cmd
}

func NewGatewayCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "gateway",
		Short: "gateway starts api gateway server.",
		Long: `
			gateway starts api gateway server.

      Find more information at:
            https://hogehoge.com`,
		Run: runHelp,
	}

	ioStreams := IOStreams{In: in, Out: out, ErrOut: err}

	cmds.AddCommand(NewCmdStart(nil, ioStreams))

	return cmds
}

type defaultPluginHandler struct{}

// Lookup implements PluginHandler
func (h *defaultPluginHandler) Lookup(filename string) (string, error) {
	// if on Windows, append the "exe" extension
	// to the filename that we are looking up.
	if runtime.GOOS == "windows" {
		filename = filename + ".exe"
	}

	return exec.LookPath(filename)
}

// Execute implements PluginHandler
func (h *defaultPluginHandler) Execute(executablePath string, cmdArgs, environment []string) error {
	return syscall.Exec(executablePath, cmdArgs, environment)
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
