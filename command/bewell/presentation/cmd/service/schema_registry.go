package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.slade360emr.com/go/base/command/bewell/application/dto"
	"gitlab.slade360emr.com/go/base/command/bewell/usecases"
)

var registry usecases.SchemaRegistryUsecase

func init() {
	// Publish schema flags
	pushSchemaCmd.Flags().String("dir", ".", "The directory containing the graphql schema files")
	pushSchemaCmd.Flags().String("file-extension", "graphql", "The extension for graphql schema files in directory")

	// Validate schema flags
	validateSchemaCmd.Flags().String("dir", ".", "The directory containing the graphql schema files")
	validateSchemaCmd.Flags().String("file-extension", "graphql", "The extension for graphql schema files in directory")

	s, err := usecases.NewSchemaRegistryImpl()
	if err != nil {
		fmt.Printf("internal error")
		os.Exit(1)
	}
	registry = s

}

// pushSchemaCmd is the command used to push/update the graphql schema for a service
var pushSchemaCmd = &cobra.Command{
	Use:     "push-schema",
	Short:   "It is used to create or update the graphql schema for a service on schema registry",
	Long:    ``, //long description
	Example: "",
	Hidden:  false,
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, _ := cmd.Flags().GetString("name")
		serviceURL, _ := cmd.Flags().GetString("url")
		version, _ := cmd.Flags().GetString("version")
		schemaDir, _ := cmd.Flags().GetString("dir")
		schemaExtensionName, _ := cmd.Flags().GetString("file-extension")
		ctx := context.Background()

		service := dto.Service{
			Name:    serviceName,
			URL:     serviceURL,
			Version: version,
		}

		status, err := registry.PublishSchema(ctx, service, schemaDir, schemaExtensionName)
		if err != nil {
			fmt.Printf("error pushing schema: %v \n", err)
			os.Exit(1)
		}

		if !status.Valid {
			fmt.Printf("Schema for %v version:%v has not been published 🤡\nMessage: %v\n", service.Name, version, status.Message)
			os.Exit(1)
		}

		fmt.Printf("Schema for service: %v version: %v successfully published 🥳\n", service, version)
		os.Exit(0)
	},
}

// validateSchemaCmd is the command used to validate the graphql schema for a service
var validateSchemaCmd = &cobra.Command{
	Use:     "validate-schema",
	Short:   "It is used to validate a service's graphql schema against the schema registry",
	Long:    ``, //long description
	Example: "",
	Hidden:  false,
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, _ := cmd.Flags().GetString("name")
		serviceURL, _ := cmd.Flags().GetString("url")
		version, _ := cmd.Flags().GetString("version")
		schemaDir, _ := cmd.Flags().GetString("dir")
		schemaExtensionName, _ := cmd.Flags().GetString("file-extension")

		ctx := context.Background()

		service := dto.Service{
			Name:    serviceName,
			URL:     serviceURL,
			Version: version,
		}

		status, err := registry.ValidateSchema(ctx, service, schemaDir, schemaExtensionName)
		if err != nil {
			fmt.Printf("error validating schema: %v \n", err)
			os.Exit(1)
		}

		if !status.Valid {
			fmt.Printf("Schema for %v version:%v is Invalid 🤡\nMessage: %v\n", service.Name, version, status.Message)
			os.Exit(1)
		}

		fmt.Printf("Schema for service: %v version: %v is Valid 🥳\n", service, version)
		os.Exit(0)
	},
}
