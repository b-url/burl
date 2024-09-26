package main

import (
	"fmt"

	"github.com/b-url/burl/cmd/burl/config"
	"github.com/b-url/burl/cmd/burl/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	ArgsAPIURL = config.APIURLName
)

type RootCommand struct {
	command *cobra.Command
}

func NewRootCommand() *RootCommand {
	cobra.OnInitialize(config.Init)

	cmd := &RootCommand{}
	cmd.command = &cobra.Command{
		Use:   "burl",
		Short: "b(ookmark)url is a command-line bookmark manager",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			bindFlags(cmd, viper.GetViper())
		},
		RunE: cmd.Execute,
	}

	cmd.command.AddCommand(NewConfigCommand().command)

	pflags := cmd.command.PersistentFlags()

	pflags.String(ArgsAPIURL, "", "API URL")
	_ = viper.BindPFlag(ArgsAPIURL, pflags.Lookup(ArgsAPIURL))

	return cmd
}

func (c *RootCommand) Execute(_ *cobra.Command, _ []string) error {
	config, err := config.New()
	if err != nil {
		return err
	}

	fmt.Println("API URL:", config.APIURL)
	// TODO: Init client with API URL and inject in tui.New()

	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	_, err = p.Run()
	return err
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable).
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name

		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
