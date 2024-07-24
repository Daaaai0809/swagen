package generate_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	generate "github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/methods"
)

const expectedYaml = `get:
  operationId: OperationID
  summary: Summary
  description: Description
  tags:
  - tag1
  - tag2
  security:
  - bearer: []
  parameters:
  - name: param1
    in: query
    description: Parameter 1
    required: true
    schema:
      type: string
      format: format
  responses:
    "200":
      description: OK
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Response'
			type: object
			properties:
			  property1:
			    type: string
			required:
			- property1
`

func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment")

	os.Setenv("PATH_DIR", "test/paths")
	os.Setenv("SCHEMA_DIR", "test/schemas")
	os.Setenv("MODEL_DIR", "test/models")
}

func TestGetPathSchema_ToYaml(t *testing.T) {
	p := &methods.GetPathSchema{
		OperationID: "OperationID",
		Summary:     "Summary",
		Description: "Description",
		Tags:        []string{"tag1", "tag2"},
		Security: []generate.Security{
			{
				"bearer": {},
			},
		},
		Parameters: generate.Parameters{
			&generate.Parameter{
				Name:        "param1",
				In:          "query",
				Description: "Parameter 1",
				Required:    true,
				Schema: generate.ParameterSchema{
					"type":   "string",
					"format": "format",
				},
			},
		},
		Responses: generate.Responses{
			"200": generate.Response{
				Description: "OK",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{
							&generate.RefSchema{
								Ref: "#/components/schemas/Response",
							},
							&generate.Schema{
								Type: "object",
								Properties: map[string]generate.Schema{
									"property1": {
										Type: "string",
									},
								},
								Required: []string{"property1"},
							},
						},
					},
				},
			},
		},
	}

	yaml, err := p.ToYaml()
	if err != nil {
		t.Errorf("ToYaml() returned an error: %v", err)
	}
	assert.Equal(t, expectedYaml, yaml)
}

func TestGenerateYamlFile(t *testing.T) {
	p := &methods.GetPathSchema{
		OperationID: "OperationID",
		Summary:     "Summary",
		Description: "Description",
		Tags:        []string{"tag1", "tag2"},
		Security: []generate.Security{
			{
				"bearer": {},
			},
		},
		Parameters: generate.Parameters{
			&generate.Parameter{
				Name:        "param1",
				In:          "query",
				Description: "Parameter 1",
				Required:    true,
				Schema: generate.ParameterSchema{
					"Type":   "string",
					"Format": "format",
				},
			},
		},
		Responses: generate.Responses{
			"200": generate.Response{
				Description: "OK",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{
							&generate.RefSchema{
								Ref: "#/components/schemas/Response",
							},
							&generate.Schema{
								Type: "object",
								Properties: map[string]generate.Schema{
									"property1": {
										Type: "string",
									},
								},
								Required: []string{"property1"},
							},
						},
					},
				},
			},
		},
	}

	path := "test"
	filename := "test.yaml"

	err := generate.GeneratePathYamlFile(p, path, filename)
	if err != nil {
		t.Errorf("GeneratePathYamlFile() returned an error: %v", err)
	}

	// move to the test directory
	err = os.Chdir(path)
	if err != nil {
		t.Errorf("Failed to change to the test directory: %v", err)
	}

	// Verify that the file was created
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("GeneratePathYamlFile() did not create the expected file: %s", filename)
	}

	// Clean up the test file
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to clean up the test file: %v", err)
	}

	err = os.Chdir("..")
	if err != nil {
		t.Errorf("Failed to change to the parent directory: %v", err)
	}

	err = os.Remove(path)
	if err != nil {
		t.Errorf("Failed to clean up the test directory: %v", err)
	}
}

func TestGenerateYamlFileWithInvalidPath(t *testing.T) {
	p := &methods.GetPathSchema{
		OperationID: "OperationID",
		Summary:     "Summary",
		Description: "Description",
		Tags:        []string{"tag1", "tag2"},
		Security: []generate.Security{
			{
				"bearer": {},
			},
		},
		Parameters: generate.Parameters{
			&generate.Parameter{
				Name:        "param1",
				In:          "query",
				Description: "Parameter 1",
				Required:    true,
				Schema: generate.ParameterSchema{
					"type":   "string",
					"format": "format",
				},
			},
		},
		Responses: generate.Responses{
			"200": generate.Response{
				Description: "OK",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{
							&generate.RefSchema{
								Ref: "#/components/schemas/Response",
							},
							&generate.Schema{
								Type: "object",
								Properties: map[string]generate.Schema{
									"property1": {
										Type: "string",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	path := "/test"
	filename := "test.yaml"

	err := generate.GeneratePathYamlFile(p, path, filename)
	if err == nil {
		t.Error("GeneratePathYamlFile() did not return an error for an invalid path")
	}
}
