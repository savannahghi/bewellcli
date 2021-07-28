package usecases

import (
	"context"

	"github.com/savannahghi/bewellcli/application/dto"
)

// Usecases is a collection of all usecases
type Usecases interface {
	SchemaRegistryUsecase
}

// NewUsecaseImpl initializes a new implementation of usecases interface
func NewUsecaseImpl() (Usecases, error) {
	registry, err := NewSchemaRegistryImpl()
	if err != nil {
		return nil, err
	}

	u := UsecaseImpl{
		registry: registry,
	}

	return &u, nil
}

// UsecaseImpl is the implementation of the usecase interface
type UsecaseImpl struct {
	registry SchemaRegistryUsecase
}

// PublishSchema pushes the schema to schema registry
func (u UsecaseImpl) PublishSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	return u.registry.PublishSchema(ctx, service, dir, extension)
}

// ValidateSchema is validates a services schema against the graph in schema registry
func (u UsecaseImpl) ValidateSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	return u.registry.ValidateSchema(ctx, service, dir, extension)
}
