package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/grafana/terraform-provider-grafana/pkg/provider"
)

func main() {
	p := provider.Provider("genimports") // Instantiate the provider so that all resources are registered
	_ = p

	if err := generateImportFiles(os.Args[1]); err != nil {
		panic(err)
	}
}

// GenerateImportFiles generates import files for all resources that use a helper defined in this package
func generateImportFiles(path string) error {
	for _, r := range provider.Resources() {
		resourcePath := filepath.Join(path, "resources", r.Name, "import.sh")
		if err := os.RemoveAll(resourcePath); err != nil { // Remove the file if it exists
			return err
		}

		if r.IDType == nil {
			log.Printf("Skipping import file generation for %s because it does not have an ID type\n", r.Name)
			continue
		}

		log.Printf("Generating import file for %s (writing to %s)\n", r.Name, resourcePath)
		if err := os.WriteFile(resourcePath, []byte(r.ImportExample()), 0600); err != nil {
			return err
		}
	}
	return nil
}
