package models_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/models"
)

const SIMPLE_TEST_MODEL_FILE_PATH = "tests/generate/models/SimpleTestModel.yaml"
const COMPLEX_TEST_MODEL_FILE_PATH = "tests/generate/models/ComplexTestModel.yaml"

var (
	expectedYaml_Simple_Model  string
	expectedYaml_Complex_Model string
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

	rootDir := currentDir[:len(currentDir)-len("generate/models")]

	// NOTE: get TestFiles from root/tests/generate/models
	expectedYaml_Simple_Model_Byte, err := os.ReadFile(fmt.Sprint(rootDir, SIMPLE_TEST_MODEL_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Complex_Model_Byte, err := os.ReadFile(fmt.Sprint(rootDir, COMPLEX_TEST_MODEL_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Simple_Model = string(expectedYaml_Simple_Model_Byte)
	expectedYaml_Complex_Model = string(expectedYaml_Complex_Model_Byte)

	os.Exit(m.Run())
}

func TestModelSchema_ToYaml_Simple_Model(t *testing.T) {
	modelTitle := "SimpleTestModel"
	modelType := "object"
	modelProperties := &generate.PropertiesMap{
		"name": &generate.Schema{
			Type: "string",
		},
		"age": &generate.Schema{
			Type:   "integer",
			Format: "int32",
		},
		"is_student": &generate.Schema{
			Type:     "boolean",
			Nullable: true,
		},
	}

	model := models.NewModelSchema(modelTitle, modelType, modelProperties)

	yaml, err := model.ToYaml()

	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Simple_Model, yaml)
}

func TestModelSchema_ToYaml_Complex_Model(t *testing.T) {
	modelTitle := "ComplexTestModel"
	modelType := "object"
	modelProperties := &generate.PropertiesMap{
		"name": &generate.Schema{
			Type: "string",
		},
		"object1": &generate.Schema{
			Type: "object",
			Properties: generate.PropertiesMap{
				"property1": &generate.Schema{
					Type: "string",
				},
				"property2": &generate.Schema{
					Type:   "string",
					Format: "date-time",
				},
				"property3": &generate.Schema{
					Type:   "number",
					Format: "double",
				},
				"property4": &generate.Schema{
					Type:     "integer",
					Format:   "int32",
					Nullable: true,
				},
			},
		},
		"array1": &generate.Schema{
			Type: "array",
			Items: &generate.Schema{
				Type: "object",
				Properties: generate.PropertiesMap{
					"array_prop1": &generate.Schema{
						Type: "string",
					},
					"array_prop2": &generate.Schema{
						Type:   "string",
						Format: "password",
					},
					"array_prop3": &generate.Schema{
						Type:   "number",
						Format: "double",
					},
					"array_prop4": &generate.Schema{
						Type:     "integer",
						Format:   "int32",
						Nullable: true,
					},
				},
			},
		},
	}

	model := models.NewModelSchema(modelTitle, modelType, modelProperties)

	yaml, err := model.ToYaml()

	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Complex_Model, yaml)
}
