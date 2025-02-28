package methods_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/methods"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment")

	os.Setenv("PATH_DIR", "test/paths")
	os.Setenv("SCHEMA_DIR", "test/schemas")
	os.Setenv("MODEL_DIR", "test/models")

	os.Exit(m.Run())
}

func TestGetPathSchema_ToYaml(t *testing.T) {
	getPathSchema := methods.NewGetPathSchema("testGet", "testGet", "Test GET endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}}, generate.Parameters{&generate.Parameter{
		Name:        "id",
		In:          "path",
		Description: "ID",
		Required:    true,
		Schema: generate.ParameterSchema{
			"type":   "integer",
			"format": "int64",
		},
	}, &generate.RefParameter{
		Ref: "#/components/parameters/QueryParameter",
	}}, generate.Responses{
		"200": {
			Description: "Success",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/SuccessResponse"}},
				},
			},
		},
		"400": {
			Description: "Bad Request",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/BadRequestResponse"}},
				},
			},
		},
		"401": {
			Description: "Unauthorized",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/UnauthorizedResponse"}},
				},
			},
		},
		"403": {
			Description: "Forbidden",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/ForbiddenResponse"}},
				},
			},
		},
		"404": {
			Description: "Not Found",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/NotFoundResponse"}},
				},
			},
		},
		"500": {
			Description: "Internal Server Error",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
				},
			},
		},
	})

	yaml, err := getPathSchema.ToYaml()
	assert.Nil(t, err)
	
	tempPath := os.TempDir()
	tempFileName := "get_test.yaml"

	err = generate.GenerateYamlFile(getPathSchema, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err_ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err_)
	assert.Equal(t, yaml, string(file))
}

func TestPostPathSchema_ToYaml(t *testing.T) {
	postPathSchema := methods.NewPostPathSchema("testPost", "testPost", "Test POST endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/RequestBody"}},
				},
			},
		},
		generate.Parameters{
			&generate.Parameter{
				Name:        "id",
				In:          "path",
				Description: "ID",
				Required:    true,
				Schema: generate.ParameterSchema{
					"type":   "integer",
					"format": "int64",
				},
			},
			&generate.RefParameter{
				Ref: "#/components/parameters/QueryParameter",
			},
		},
		generate.Responses{
			"201": {
				Description: "Created",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/CreatedResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := postPathSchema.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "post_test.yaml"

	err = generate.GenerateYamlFile(postPathSchema, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err_ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err_)
	assert.Equal(t, yaml, string(file))
}

func TestPutPathSchema_ToYaml(t *testing.T) {
	putPathSchema := methods.NewPutPathSchema("testPut", "testPut", "Test PUT endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/RequestBody"}},
				},
			},
		},
		generate.Parameters{
			&generate.Parameter{
				Name:        "id",
				In:          "path",
				Description: "ID",
				Required:    true,
				Schema: generate.ParameterSchema{
					"type":   "integer",
					"format": "int64",
				},
			},
			&generate.RefParameter{
				Ref: "#/components/parameters/QueryParameter",
			},
		},
		generate.Responses{
			"200": {
				Description: "Success",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/SuccessResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := putPathSchema.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "put_test.yaml"

	err = generate.GenerateYamlFile(putPathSchema, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err_ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err_)
	assert.Equal(t, yaml, string(file))
}

func TestDeletePathSchema_ToYaml(t *testing.T) {
	deletePathSchema := methods.NewDeletePathSchema("testDelete", "testDelete", "Test DELETE endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/RequestBody"}},
				},
			},
		},
		generate.Parameters{
			&generate.Parameter{
				Name:        "id",
				In:          "path",
				Description: "ID",
				Required:    true,
				Schema: generate.ParameterSchema{
					"type":   "integer",
					"format": "int64",
				},
			},
			&generate.RefParameter{
				Ref: "#/components/parameters/QueryParameter",
			},
		},
		generate.Responses{
			"200": {
				Description: "Success",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/SuccessResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.Schema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := deletePathSchema.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "delete_test.yaml"

	err = generate.GenerateYamlFile(deletePathSchema, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err_ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err_)
	assert.Equal(t, yaml, string(file))
}
