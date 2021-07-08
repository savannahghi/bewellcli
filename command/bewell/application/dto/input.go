package dto

import "fmt"

// Service hold the information about a service
// This information is obtained from the command line
// and represents details of the service being worked on
type Service struct {
	Name    string
	URL     string
	Version string
}

// ValidateFields ensures service related fields are not empty
func (s Service) ValidateFields() error {
	if s.Name == "" || s.Version == "" || s.URL == "" {
		return fmt.Errorf("missing required service details. required flags: --name, --version and --url")
	}

	return nil
}
