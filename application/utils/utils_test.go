package utils_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/savannahghi/bewellcli/application/utils"
	"github.com/savannahghi/bewellcli/domain"
)

// CoverageThreshold sets the test coverage threshold below which the tests will fail
// const CoverageThreshold = 0

// func TestMain(m *testing.M) {
// 	existingDebug, err := serverutils.GetEnvVar("DEBUG")
// 	if err != nil {
// 		existingDebug = "false"
// 	}

// 	os.Setenv("DEBUG", "true")

// 	rc := m.Run()
// 	// Restore DEBUG envar to original value after running test
// 	os.Setenv("DEBUG", existingDebug)

// 	// rc 0 means we've passed,
// 	// and CoverMode will be non empty if run with -cover
// 	if rc == 0 && testing.CoverMode() != "" {
// 		c := testing.Coverage()
// 		if c < CoverageThreshold {
// 			fmt.Println("Tests passed but coverage failed at", c)
// 			rc = -1
// 		}
// 	}

// 	os.Exit(rc)
// }

func Test_ReadSchemaFilesInDirectory(t *testing.T) {
	mutation := `
	type Mutation {
		world: String
	  }
	`
	query := `
	type Query {
		world: String
	  }
	`

	combinedSchema := mutation + "\n" + query + "\n"

	testDir1 := t.TempDir()

	schemaFile1 := filepath.Join(testDir1, "mutation.graphql")
	schemaFile2 := filepath.Join(testDir1, "query.graphql")

	err := ioutil.WriteFile(schemaFile1, []byte(mutation), 0666)
	if err != nil {
		t.Errorf("error writing to test file: %v", err)
		return
	}
	err = ioutil.WriteFile(schemaFile2, []byte(query), 0666)
	if err != nil {
		t.Errorf("error writing to test file: %v", err)
		return
	}

	testDir2 := t.TempDir()

	type args struct {
		dir       string
		extension string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success:directory with schema files",
			args: args{
				dir:       testDir1,
				extension: "graphql",
			},
			want:    combinedSchema,
			wantErr: false,
		},
		{
			name: "fail:directory without schema files",
			args: args{
				dir:       testDir2,
				extension: "graphql",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "success:single schema file",
			args: args{
				dir:       schemaFile1,
				extension: "graphql",
			},
			want:    mutation + "\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.ReadSchemaFilesInDirectory(tt.args.dir, tt.args.extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSchemaFilesInDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadSchemaFilesInDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaRegistryRequest(t *testing.T) {

	schema := `
	type Query {
		world: String
	  }
	`
	url := "https://postman-echo.com/post"
	invalidUrl := "http://example.com"

	payload := domain.GraphqlSchemaPayload{
		Name:     "test",
		URL:      url,
		Version:  "0.0.1",
		TypeDefs: schema,
	}

	invalidPayload := domain.GraphqlSchemaPayload{
		Name:     "test",
		URL:      invalidUrl,
		Version:  "0.0.1",
		TypeDefs: schema,
	}

	type args struct {
		payload domain.GraphqlSchemaPayload
		url     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case :) schema registry request",
			args: args{
				payload: payload,
				url:     url,
			},
			wantErr: false,
		},

		{
			name: "sad case :( invalid payload",
			args: args{
				payload: invalidPayload,
				url:     invalidUrl,
			},
			wantErr: true,
		},
		{
			name: "sad case :( missing url",
			args: args{
				payload: invalidPayload,
				url:     "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		_, err := utils.SchemaRegistryRequest(tt.args.payload, tt.args.url)
		if err != nil && !tt.wantErr {
			t.Errorf("did not get expected error but got: %v", err)
		}
	}

}
