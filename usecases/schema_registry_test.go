package usecases_test

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/savannahghi/bewellcli/application/dto"
	"github.com/savannahghi/bewellcli/usecases"
)

func TestMain(m *testing.M) {
	// Test url
	os.Setenv("REGISTRY_URL", "https://postman-echo.com/post")

	rc := m.Run()

	os.Exit(rc)
}

func Test_validateSchema(t *testing.T) {
	s, err := usecases.NewSchemaRegistryImpl()
	if err != nil {
		t.Errorf("cannot initialize schema registry implementation")
		return
	}
	schema := `
	type Query {
		world: String
	  }
	`
	testDir := t.TempDir()
	schemaFile := filepath.Join(testDir, "test.graphql")
	err = ioutil.WriteFile(schemaFile, []byte(schema), 0666)
	if err != nil {
		t.Errorf("error writing to test file: %v", err)
		return
	}

	emptyTestDir := t.TempDir()

	testService := dto.Service{
		Name:    "bewellcli",
		URL:     "https://bewell-test.com",
		Version: "0.0.1",
	}

	type args struct {
		ctx       context.Context
		service   dto.Service
		dir       string
		extension string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: validation request",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       testDir,
				extension: "graphql",
			},
			wantErr: true,
		},
		{
			name: "fail: no schema files in directory",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       emptyTestDir,
				extension: "graphql",
			},
			wantErr: true,
		},
		{
			name: "fail: missing url",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       testDir,
				extension: "graphql",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := s.ValidateSchema(tt.args.ctx, tt.args.service, tt.args.dir, tt.args.extension); (err != nil) != tt.wantErr {
				t.Errorf("validateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pushSchema(t *testing.T) {
	s, err := usecases.NewSchemaRegistryImpl()
	if err != nil {
		t.Errorf("cannot initialize schema registry implementation")
		return
	}

	schema := `
	type Query {
		world: String
	  }
	`
	testDir := t.TempDir()
	schemaFile := filepath.Join(testDir, "test.graphql")
	err = ioutil.WriteFile(schemaFile, []byte(schema), 0666)
	if err != nil {
		t.Errorf("error writing to test file: %v", err)
		return
	}

	emptyTestDir := t.TempDir()

	testService := dto.Service{
		Name:    "bewellcli",
		URL:     "https://bewell-test.com",
		Version: "0.0.1",
	}

	// TODO: Update test url
	testValidationURL := "https://postman-echo.com/post"

	type args struct {
		ctx       context.Context
		service   dto.Service
		dir       string
		extension string
		pushURL   string
	}
	tests := []struct {
		name    string
		args    args
		want    *dto.SchemaStatus
		wantErr bool
	}{
		{
			name: "success: push request",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       testDir,
				extension: "graphql",
				pushURL:   testValidationURL,
			},
			wantErr: true,
			want: &dto.SchemaStatus{
				Valid: false,
			},
		},
		{
			name: "fail: no schema files in directory",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       emptyTestDir,
				extension: "graphql",
				pushURL:   testValidationURL,
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "fail: missing url",
			args: args{
				ctx:       context.Background(),
				service:   testService,
				dir:       testDir,
				extension: "graphql",
				pushURL:   "",
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.PublishSchema(tt.args.ctx, tt.args.service, tt.args.dir, tt.args.extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("publishSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
