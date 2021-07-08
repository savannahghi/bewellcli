package usecases

import (
	"context"
	"net/url"

	"gitlab.slade360emr.com/go/base"
	"gitlab.slade360emr.com/go/base/command/bewell/application/dto"
	"gitlab.slade360emr.com/go/base/command/bewell/application/utils"
	"gitlab.slade360emr.com/go/base/command/bewell/domain"
)

// SchemaRegistry variables
var (
	publishSchemaURLPath  = "schema/push"
	validateSchemaURLPath = "schema/validate"
	schemaRegistryURLEnv  = "REGISTRY_URL"
)

// SchemaRegistryUsecase defines methods for the schema registry
type SchemaRegistryUsecase interface {
	PublishSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error)
	ValidateSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error)
}

// SchemaRegistryImpl ...
type SchemaRegistryImpl struct {
	RegistryURL *url.URL
}

// NewSchemaRegistryImpl new Schema registry usecase implementation
func NewSchemaRegistryImpl() (SchemaRegistryUsecase, error) {

	baseURL := base.MustGetEnvVar(schemaRegistryURLEnv)
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	n := SchemaRegistryImpl{
		RegistryURL: u,
	}

	return &n, nil
}

// PublishSchema pushes the schema to schema registry
func (s SchemaRegistryImpl) PublishSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	err := service.ValidateFields()
	if err != nil {
		return nil, err
	}

	schema, err := utils.ReadSchemaFilesInDirectory(dir, extension)
	if err != nil {
		return nil, err
	}

	payload := domain.GraphqlSchemaPayload{
		Name:     service.Name,
		URL:      service.URL,
		Version:  service.Version,
		TypeDefs: schema,
	}

	s.RegistryURL.Path = publishSchemaURLPath
	publishSchemaURL := s.RegistryURL.String()

	status, err := utils.SchemaRegistryRequest(payload, publishSchemaURL)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// ValidateSchema is validates a services schema against the graph in schema registry
func (s SchemaRegistryImpl) ValidateSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	err := service.ValidateFields()
	if err != nil {
		return nil, err
	}

	schema, err := utils.ReadSchemaFilesInDirectory(dir, extension)
	if err != nil {
		return nil, err
	}

	payload := domain.GraphqlSchemaPayload{
		Name:     service.Name,
		URL:      service.URL,
		Version:  service.Version,
		TypeDefs: schema,
	}

	s.RegistryURL.Path = validateSchemaURLPath
	validationURL := s.RegistryURL.String()

	status, err := utils.SchemaRegistryRequest(payload, validationURL)
	if err != nil {
		return nil, err
	}

	return status, nil
}
