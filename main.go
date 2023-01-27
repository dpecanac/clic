// Package main ...
package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := createCommand().Execute(); err != nil {
		os.Exit(0)
	}

	os.Exit(1)
}

func createCommand() *cobra.Command {
	config := &Config{}

	command := &cobra.Command{
		Use:   `clic`,
		Short: `clic is a command line api client `,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileName := DEFAULT_DATA_FILE

			if config.File != "" {
				fileName = config.File
			}

			data, err := read(fileName)
			if err != nil {
				println(err.Error())
				return err
			}

			if data == nil {
				println("Could not read data file")
				return nil
			}

			run(*data, *config)
			return nil
		},
	}

	command.Flags().StringVarP(&config.File, "file", "f", "", `path to the data file (default is "clic.yaml")`)
	command.Flags().StringVarP(&config.Name, "names", "n", "", "name of the requests to execute, if not provided all requests will be executed")
	command.Flags().StringVarP(&config.Tags, "tags", "t", "", "tags of the requests to execute, if not provided all requests will be executed")
	command.Flags().BoolVarP(&config.Verbose, "verbose", "v", false, "print response details")

	return command
}
