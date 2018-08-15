package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

type List struct {
}

type ListOptions struct {
	// flags
	OutputFileType string
	OutputFilePath string

	IOStreams
}

type Data struct {
	Events []Event `yaml:"events"`
}

type Event struct {
	Title    string     `yaml:"title"`
	Date     CustomTime `yaml:"date"`
	Category string     `yaml:"category"`
}

type CustomTime struct {
	time.Time
}

func NewListOptions(ioStreams IOStreams) *ListOptions {
	return &ListOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdList(f cmdutil.Factory, ioStreams IOStreams) *cobra.Command {
	o := NewListOptions(ioStreams)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// cmdutil.CheckErr(o.Complete(f, cmd))
			// cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.Run(args))
		},
	}
	cmd.PersistentFlags().StringVar(&o.OutputFileType, "type", "yaml", "yaml or json.")
	cmd.PersistentFlags().StringVar(&o.OutputFilePath, "out", "", "output file path.")
	cmd.MarkFlagRequired("out")
	return cmd
}

// func (o *ListOptions) Complete(f cmdutil.Factory, cmd *cobra.Command) error {
// 	var err error
// 	o.discoveryClient, err = f.ToDiscoveryClient()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (o *ListOptions) Validate() error {
// 	if o.Output != "" && o.Output != "yaml" && o.Output != "json" {
// 		return errors.New(`--output must be 'yaml' or 'json'`)
// 	}
// 	return nil
// }

func (o *ListOptions) Run(args []string) error {
	if len(args) > 1 || len(args) < 1 {
		return errors.New("input file is requred.")
	}

	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}
	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		return err
	}

	sortData(d.Events)

	if o.OutputFilePath == "" {
		o.OutputFilePath = "output"
	}
	var s []byte
	switch o.OutputFileType {
	case "yaml":
		s, err = yaml.Marshal(&d)
		if err != nil {
			return err
		}
		fmt.Println(string(s))

		if o.OutputFilePath == "" {
			o.OutputFilePath += ".yaml"
		}

	case "json":
		// TODO
		// if o.OutputFilePath == "" {
		// 	o.OutputFilePath += ".json"
		// }
		return errors.New("json is not impremented.")
	default:
		return fmt.Errorf("default", o.OutputFileType)
	}

	file, err := os.Create(o.OutputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(([]byte)(s))

	return nil
}

func (ct CustomTime) MarshalYAML() (interface{}, error) {
	return fmt.Sprintf("%d/%d/%d", ct.Year(), int(ct.Month()), ct.Day()), nil
}

func (ct *CustomTime) Format(s string) string {
	return fmt.Sprintf("%d/%d/%d", ct.Year(), int(ct.Month()), ct.Day())
}

func (ct *CustomTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	d := ""
	err := unmarshal(&d)
	if err != nil {
		return err
	}

	s := strings.Split(d, "/")

	var yi, mi, di int
	// day
	if len(s) >= 3 {
		di, err = strconv.Atoi(s[2])
		if err != nil {
			return err
		}
	}
	// month
	if len(s) >= 2 {
		mi, err = strconv.Atoi(s[1])
		if err != nil {
			return err
		}
		if mi > 12 {
			return errors.New("month should be 1-12.")
		}
	}
	// year
	yi, err = strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	// location
	utc, _ := time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	if mi == 0 {
		ct.Time = time.Date(yi, 1, 1, 0, 0, 0, 0, utc)
		return nil
	}
	if di == 0 {
		ct.Time = time.Date(yi, time.Month(mi), 1, 0, 0, 0, 0, utc)
		return nil
	}
	ct.Time = time.Date(yi, time.Month(mi), di, 0, 0, 0, 0, utc)
	return nil
}

func sortData(events []Event) {
	sort.Slice(events, func(i, j int) bool {
		if events[j].Date.Time.After(events[i].Date.Time) {
			return true
		}
		return false
	})
}
