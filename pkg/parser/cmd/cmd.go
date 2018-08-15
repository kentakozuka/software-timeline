package cmd

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
)

func NewDefaultParserCommand() *cobra.Command {
	return NewDefaultParserCommandWithArgs(&defaultPluginHandler{}, os.Args, os.Stdin, os.Stdout, os.Stderr)
}

func NewDefaultParserCommandWithArgs(pluginHandler PluginHandler, args []string, in io.Reader, out, errout io.Writer) *cobra.Command {
	cmd := NewParserCommand(in, out, errout)

	if pluginHandler == nil {
		return cmd
	}

	// if len(args) > 1 {
	// 	cmdPathPieces := args[1:]

	// 	// only look for suitable extension executables if
	// 	// the specified command does not already exist
	// 	if _, _, err := cmd.Find(cmdPathPieces); err != nil {
	// 		if err := handleEndpointExtensions(pluginHandler, cmdPathPieces); err != nil {
	// 			fmt.Fprintf(errout, "%v\n", err)
	// 			os.Exit(1)
	// 		}
	// 	}
	// }

	return cmd
}

func NewParserCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "parser",
		Short: "parser parse text file.",
		Long: `
      parser parse text file.

      Find more information at:
            https://hogehoge.com`,
		// parser itself is not executable. it diplays help if called solely.
		Run: runHelp,
	}

	// flags := cmds.PersistentFlags()
	// flags.SetNormalizeFunc(utilflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	// flags.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)

	// kubeConfigFlags := genericclioptions.NewConfigFlags()
	// kubeConfigFlags.AddFlags(flags)
	// matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	// matchVersionKubeConfigFlags.AddFlags(cmds.PersistentFlags())

	// cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// f := cmdutil.NewFactory(matchVersionKubeConfigFlags)

	// From this point and forward we get warnings on flags that contain "_" separators
	// cmds.SetGlobalNormalizationFunc(utilflag.WarnWordSepNormalizeFunc)

	ioStreams := IOStreams{In: in, Out: out, ErrOut: err}

	// groups := templates.CommandGroups{
	// 	{
	// 		Message: "Deploy Commands:",
	// 		Commands: []*cobra.Command{
	// 			rollout.NewCmdRollout(f, ioStreams),
	// 			NewCmdRollingUpdate(f, ioStreams),
	// 			NewCmdScale(f, ioStreams),
	// 			NewCmdAutoscale(f, ioStreams),
	// 		},
	// 	},
	// }
	// groups.Add(cmds)

	// filters := []string{"options"}

	// // Hide the "alpha" subcommand if there are no alpha commands in this build.
	// alpha := NewCmdAlpha(f, ioStreams)
	// if !alpha.HasSubCommands() {
	// 	filters = append(filters, alpha.Name())
	// }

	// templates.ActsAsRootCommand(cmds, filters, groups...)

	// for name, completion := range bash_completion_flags {
	// 	if cmds.Flag(name) != nil {
	// 		if cmds.Flag(name).Annotations == nil {
	// 			cmds.Flag(name).Annotations = map[string][]string{}
	// 		}
	// 		cmds.Flag(name).Annotations[cobra.BashCompCustom] = append(
	// 			cmds.Flag(name).Annotations[cobra.BashCompCustom],
	// 			completion,
	// 		)
	// 	}
	// }

	// cmds.AddCommand(cmdconfig.NewCmdConfig(f, clientcmd.NewDefaultPathOptions(), ioStreams))
	// cmds.AddCommand(NewCmdVersion(f, ioStreams))

	cmds.AddCommand(NewCmdList(nil, ioStreams))
	// cmds.AddCommand(NewCmdList(f, ioStreams))

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
