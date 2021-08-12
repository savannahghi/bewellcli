package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/savannahghi/bewellcli/domain"
	"github.com/savannahghi/serverutils"
)

// CoverageThreshold sets the test coverage threshold below which the tests will fail
const CoverageThreshold = 0.71

func TestMain(m *testing.M) {
	os.Setenv("MESSAGE_KEY", "this-is-a-test-key$$$")
	os.Setenv("ENVIRONMENT", "staging")
	err := os.Setenv("ROOT_COLLECTION_SUFFIX", "staging")
	if err != nil {
		if serverutils.IsDebug() {
			log.Printf("can't set root collection suffix in env: %s", err)
		}
		os.Exit(-1)
	}
	existingDebug, err := serverutils.GetEnvVar("DEBUG")
	if err != nil {
		existingDebug = "false"
	}

	os.Setenv("DEBUG", "true")

	rc := m.Run()
	// Restore DEBUG envar to original value after running test
	os.Setenv("DEBUG", existingDebug)

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < CoverageThreshold {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}

	os.Exit(rc)
}

func Test_readSchemaFile(t *testing.T) {
	schema := `
	type Query {
		world: String
	  }
	`
	testDir := t.TempDir()
	schemaFile := filepath.Join(testDir, "test.graphql")
	err := ioutil.WriteFile(schemaFile, []byte(schema), 0666)
	if err != nil {
		t.Errorf("error writing to test file: %v", err)
		return
	}

	type args struct {
		schemaFile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success:existing test file",
			args: args{
				schemaFile: schemaFile,
			},
			want:    schema,
			wantErr: false,
		},
		{
			name: "fail:missing file",
			args: args{
				schemaFile: "doesn't exist",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSchemaFile(tt.args.schemaFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("readSchemaFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readSchemaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Change placeholder request for testing
func Test_makeRequest(t *testing.T) {
	schema := `
	type Query {
		world: String
	  }
	`
	// TODO: Update the test URL
	url := "https://postman-echo.com/post"

	type args struct {
		url   string
		body  domain.GraphqlSchemaPayload
		debug bool
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "success: debug true",
			args: args{
				url: url,
				body: domain.GraphqlSchemaPayload{
					Name:     "test",
					Version:  "0.0.1",
					TypeDefs: schema,
				},
				debug: true,
			},
			wantErr: false,
		},
		{
			name: "success: make request",
			args: args{
				url: url,
				body: domain.GraphqlSchemaPayload{
					Name:     "test",
					Version:  "0.0.1",
					TypeDefs: schema,
				},
				debug: false,
			},
			wantErr: false,
		},
		{
			name: "fail: empty url in request",
			args: args{
				url:   "",
				body:  domain.GraphqlSchemaPayload{},
				debug: false,
			},
			wantErr: true,
		},
		{
			name: "success: make request",
			args: args{
				url: url,
				body: domain.GraphqlSchemaPayload{
					Name:     "test",
					Version:  "0.0.1",
					TypeDefs: "",
				},
				debug: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// set debug environment
		os.Setenv("DEBUG", strconv.FormatBool(tt.args.debug))
		t.Run(tt.name, func(t *testing.T) {
			_, err := makeRequest(tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	// Reset debug environment
	os.Setenv("DEBUG", "false")
}
