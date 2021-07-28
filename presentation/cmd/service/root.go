package cmd

import (
	"github.com/spf13/cobra"
)

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Service deals with Bewell microservices",
	Long:  ``,
}

func init() {
	// Add sub commands for service
	ServiceCmd.AddCommand(validateSchemaCmd)
	ServiceCmd.AddCommand(pushSchemaCmd)
	ServiceCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Persistent Flags which will work for this command and all subcommands
	// e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")
	ServiceCmd.PersistentFlags().StringP("version", "v", "", "Version of a service")
	ServiceCmd.PersistentFlags().StringP("name", "n", "", "The name of a service")
	ServiceCmd.PersistentFlags().StringP("url", "u", "", "The domain url for the service")
}
