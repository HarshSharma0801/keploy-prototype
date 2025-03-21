package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3" 
)

func main() {
	var rootCmd = &cobra.Command{Use: "keploy", Short: "Keploy Contract Testing Tool"}
	rootCmd.AddCommand(generateCmd(), validateCmd(), downloadCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func generateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate and save OpenAPI schemas from HTTPDoc",
		Run: func(cmd *cobra.Command, args []string) {
			tests := loadSampleTests()
			mocks := loadSampleMocks()
			
			// Save tests to provider directory
			for name, doc := range tests {
				schema := HTTPDocToOpenAPI(doc)
				saveSchema(schema, "ecom-service/v1/tests/contracts/provider", name+".yaml")
				fmt.Printf("Saved schema for test %s\n", name)
			}
			
			// Save mocks to consumer directory
			for name, doc := range mocks {
				schema := HTTPDocToOpenAPI(doc)
				saveSchema(schema, "ecom-service/v1/tests/contracts/consumer", name+".yaml")
				fmt.Printf("Saved schema for mock %s\n", name)
			}
		},
	}
}

func validateCmd() *cobra.Command {
	var mode string
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate contracts (consumer or provider mode)",
		Run: func(cmd *cobra.Command, args []string) {
			tests := convertDocsToOpenAPI(loadSampleTests())
			mocks := convertDocsToOpenAPI(loadSampleMocks())
			if mode == "consumer" {
				fmt.Println("Running Consumer-Driven Validation:")
				ValidateConsumer(tests, mocks)
			} else if mode == "provider" {
				fmt.Println("Running Provider-Driven Validation:")
				ValidateProvider(tests, mocks)
			} else {
				fmt.Println("Invalid mode. Use 'consumer' or 'provider'.")
			}
		},
	}
	cmd.Flags().StringVarP(&mode, "mode", "m", "consumer", "Validation mode: consumer or provider")
	return cmd
}

func downloadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "Download contract artifacts",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading artifacts... (Not implemented)")
		},
	}
}

func convertDocsToOpenAPI(docs map[string]HTTPDoc) map[string]OpenAPI {
	result := make(map[string]OpenAPI)
	for name, doc := range docs {
		result[name] = HTTPDocToOpenAPI(doc)
	}
	return result
}

// saveSchema saves an OpenAPI schema to a file in YAML format
func saveSchema(schema OpenAPI, dir, filename string) {
	data, err := yaml.Marshal(schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling schema: %v\n", err)
		return
	}
	
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", dir, err)
		return
	}
	
	// Write file
	path := filepath.Join(dir, filename)
	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", path, err)
	}
}