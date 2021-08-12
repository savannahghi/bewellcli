package usecases_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/savannahghi/bewellcli/application/dto"
	"github.com/savannahghi/bewellcli/usecases"
	"github.com/stretchr/testify/assert"
)

type FakeSchemaRegistryImpl struct {
	RegistryURL *url.URL
}

func NewFakeSchemaRegistryImpl() (*FakeSchemaRegistryImpl, error) {
	baseURL := "https://postman-echo.com/post"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	n := FakeSchemaRegistryImpl{
		RegistryURL: u,
	}

	return &n, nil
}

func (s FakeSchemaRegistryImpl) PublishSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	status := &dto.SchemaStatus{}
	err := fmt.Errorf("error example")
	if ctx == nil {
		return nil, err
	}
	return status, nil
}

func (s FakeSchemaRegistryImpl) ValidateSchema(ctx context.Context, service dto.Service, dir, extension string) (*dto.SchemaStatus, error) {
	status := &dto.SchemaStatus{}
	err := fmt.Errorf("error example")
	if ctx == nil {
		return nil, err
	}
	return status, nil
}

// func initializeFakeUsecases(haserr bool) (*usecases.Usecases, error) {
// 	registry, _ := usecases.NewUsecaseImpl()
// 	_, err := initializeFakeNewSchemaRegistryImpl(haserr)
// 	if err != nil {
// 		return nil, fmt.Errorf("has errors: %s", err)
// 	}
// 	return &registry, nil
// }

// func initializeFakeNewSchemaRegistryImpl(haserr bool) (usecases.SchemaRegistryUsecase, error) {
// 	if haserr {
// 		return nil, fmt.Errorf("has err")
// 	}
// 	n := usecases.SchemaRegistryImpl{}
// 	return &n, nil
// }
func TestNewUsecaseImpl(t *testing.T) {
	ctx := context.Background()
	service := dto.Service{}
	dir := "https://postman-echo.com/post"
	extension := "graphql"

	registry, err := usecases.NewUsecaseImpl()
	assert.NotNil(t, registry)
	assert.Nil(t, err)

	// fakeusecase, err := NewFakeSchemaRegistryImpl()
	// assert.Nil(t, err)
	// assert.NotNil(t, fakeusecase)

	registry.PublishSchema(ctx, service, dir, extension)
	// assert.Nil(t, err)
	// assert.NotNil(t, status)
	registry.ValidateSchema(ctx, service, dir, extension)
	// assert.Nil(t, err)
	// assert.NotNil(t, status)
}
