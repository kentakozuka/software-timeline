package cmd

import (
	"log"

	"github.com/spf13/cobra"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

type Start struct {
}

type StartOptions struct {
	// flags
	OutputFileType string
	OutputFilePath string

	IOStreams
}

func NewStartOptions(ioStreams IOStreams) *StartOptions {
	return &StartOptions{
		IOStreams: ioStreams,
	}
}

func NewCmdStart(f cmdutil.Factory, ioStreams IOStreams) *cobra.Command {
	o := NewStartOptions(ioStreams)
	cmd := &cobra.Command{
		Use:   "start",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Run(args))
		},
	}
	return cmd
}

func (o *StartOptions) Run(args []string) error {
	server := NewServer()

	log.Printf("Serving on https://0.0.0.0:8000")
	server.ListenAndServe()

	return nil
}
