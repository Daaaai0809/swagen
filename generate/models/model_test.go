package models_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/models"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment")

	os.Setenv("PATH_DIR", "test/paths")
	os.Setenv("SCHEMA_DIR", "test/schemas")
	os.Setenv("MODEL_DIR", "test/models")

	os.Exit(m.Run())
}

func TestModelSchema_ToYaml_Simple_Model(t *testing.T) {
	modelName := "SimpleTestModel"
	modelTitle := "SimpleTestModel"
	modelType := "object"
	modelProperties := generate.PropertiesMap{
		"name": generate.Schema{
			Type: "string",
		},
		"age": generate.Schema{
			Type:   "integer",
			Format: "int32",
		},
		"is_student": generate.Schema{
			Type:     "boolean",
			Nullable: true,
		},
	}

	model := models.NewModelSchema(modelName, modelTitle, modelType, modelProperties)

	yaml, err := model.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "simple_test_model.yaml"

	err = generate.GenerateYamlFile(model, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err)
	assert.Equal(t, yaml, string(file))
}

func TestModelSchema_ToYaml_Complex_Model(t *testing.T) {
	modelName := "ComplexTestModel"
	modelTitle := "ComplexTestModel"
	modelType := "object"
	modelProperties := generate.PropertiesMap{
		"name": generate.Schema{
			Type: "string",
		},
		"object1": generate.Schema{
			Type: "object",
			Properties: generate.PropertiesMap{
				"property1": generate.Schema{
					Type: "string",
				},
				"property2": generate.Schema{
					Type:   "string",
					Format: "date-time",
				},
				"property3": generate.Schema{
					Type:   "number",
					Format: "double",
				},
				"property4": generate.Schema{
					Type:     "integer",
					Format:   "int32",
					Nullable: true,
				},
			},
		},
		"array1": generate.Schema{
			Type: "array",
			Items: &generate.Schema{
				Type: "object",
				Properties: generate.PropertiesMap{
					"array_prop1": generate.Schema{
						Type: "string",
					},
					"array_prop2": generate.Schema{
						Type:   "string",
						Format: "password",
					},
					"array_prop3": generate.Schema{
						Type:   "number",
						Format: "double",
					},
					"array_prop4": generate.Schema{
						Type:     "integer",
						Format:   "int32",
						Nullable: true,
					},
				},
			},
		},
	}

	model := models.NewModelSchema(modelName, modelTitle, modelType, modelProperties)

	yaml, err := model.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "complex_test_model.yaml"

	err = generate.GenerateYamlFile(model, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err)
	assert.Equal(t, yaml, string(file))
}
