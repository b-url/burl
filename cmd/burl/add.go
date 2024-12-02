package main

import (
	"database/sql"

	"github.com/b-url/burl/cmd/burl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AddCommand struct {
	db      *sql.DB
	command *cobra.Command
}

func NewAddCommand() *AddCommand {
	cobra.OnInitialize(config.Init)

	cmd := &AddCommand{}
	cmd.command = &cobra.Command{
		Use:   "add",
		Short: "add bookmark",
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

func (c *AddCommand) Execute(_ *cobra.Command, _ []string) error {
	// config, err := config.New()
	// if err != nil {
	// 	return err
	// }

	// fmt.Print("hello")
	// return err
	return nil
}
