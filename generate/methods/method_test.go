package methods_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/methods"
)

const (
	GET_PATH_SCHEMA_FILE_PATH    = "tests/generate/methods/Get.yaml"
	POST_PATH_SCHEMA_FILE_PATH   = "tests/generate/methods/Post.yaml"
	PUT_PATH_SCHEMA_FILE_PATH    = "tests/generate/methods/Put.yaml"
	DELETE_PATH_SCHEMA_FILE_PATH = "tests/generate/methods/Delete.yaml"
)

var (
	expectedYaml_Get_Path_Schema    string
	expectedYaml_Post_Path_Schema   string
	expectedYaml_Put_Path_Schema    string
	expectedYaml_Delete_Path_Schema string
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment")

	os.Setenv("PATH_DIR", "test/paths")
	os.Setenv("SCHEMA_DIR", "test/schemas")
	os.Setenv("MODEL_DIR", "test/models")

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting root directory: ", err)
		os.Exit(1)
	}

	rootDir := currentDir[:len(currentDir)-len("generate/methods")]

	// NOTE: get TestFiles from root/tests/generate/methods
	expectedYaml_Get_Path_Schema_Byte, err := os.ReadFile(fmt.Sprint(rootDir, GET_PATH_SCHEMA_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Post_Path_Schema_Byte, err := os.ReadFile(fmt.Sprint(rootDir, POST_PATH_SCHEMA_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Put_Path_Schema_Byte, err := os.ReadFile(fmt.Sprint(rootDir, PUT_PATH_SCHEMA_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Delete_Path_Schema_Byte, err := os.ReadFile(fmt.Sprint(rootDir, DELETE_PATH_SCHEMA_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Get_Path_Schema = string(expectedYaml_Get_Path_Schema_Byte)
	expectedYaml_Post_Path_Schema = string(expectedYaml_Post_Path_Schema_Byte)
	expectedYaml_Put_Path_Schema = string(expectedYaml_Put_Path_Schema_Byte)
	expectedYaml_Delete_Path_Schema = string(expectedYaml_Delete_Path_Schema_Byte)

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
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/SuccessResponse"}},
				},
			},
		},
		"400": {
			Description: "Bad Request",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/BadRequestResponse"}},
				},
			},
		},
		"401": {
			Description: "Unauthorized",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/UnauthorizedResponse"}},
				},
			},
		},
		"403": {
			Description: "Forbidden",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/ForbiddenResponse"}},
				},
			},
		},
		"404": {
			Description: "Not Found",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/NotFoundResponse"}},
				},
			},
		},
		"500": {
			Description: "Internal Server Error",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
				},
			},
		},
	})

	yaml, err := getPathSchema.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Get_Path_Schema, yaml)
}

func TestPostPathSchema_ToYaml(t *testing.T) {
	postPathSchema := methods.NewPostPathSchema("testPost", "testPost", "Test POST endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/RequestBody"}},
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
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/CreatedResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := postPathSchema.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Post_Path_Schema, yaml)
}

func TestPutPathSchema_ToYaml(t *testing.T) {
	putPathSchema := methods.NewPutPathSchema("testPut", "testPut", "Test PUT endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/RequestBody"}},
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
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/SuccessResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := putPathSchema.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Put_Path_Schema, yaml)
}

func TestDeletePathSchema_ToYaml(t *testing.T) {
	deletePathSchema := methods.NewDeletePathSchema("testDelete", "testDelete", "Test DELETE endpoint", []string{"Tests"}, []generate.Security{{"Bearer": {}}},
		generate.RequestBody{
			Description: "Request body",
			Content: map[string]generate.Content{
				"application/json": {
					Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/RequestBody"}},
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
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/SuccessResponse"}},
					},
				},
			},
			"400": {
				Description: "Bad Request",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/BadRequestResponse"}},
					},
				},
			},
			"401": {
				Description: "Unauthorized",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/UnauthorizedResponse"}},
					},
				},
			},
			"403": {
				Description: "Forbidden",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/ForbiddenResponse"}},
					},
				},
			},
			"404": {
				Description: "Not Found",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/NotFoundResponse"}},
					},
				},
			},
			"500": {
				Description: "Internal Server Error",
				Content: map[string]generate.Content{
					"application/json": {
						Schema: generate.ContentSchema{&generate.RefSchema{Ref: "#/components/schemas/InternalServerErrorResponse"}},
					},
				},
			},
		})

	yaml, err := deletePathSchema.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Delete_Path_Schema, yaml)
}
