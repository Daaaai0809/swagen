package messages_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/messages"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment")

	os.Setenv("PATH_DIR", "test/paths")
	os.Setenv("SCHEMA_DIR", "test/schemas")
	os.Setenv("MODEL_DIR", "test/models")

	os.Exit(m.Run())
}

func TestMessage_ToYaml_Simple_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()

	msgType := "object"
	msgProperties["dataObject"] = generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4"},
		Properties: generate.PropertiesMap{
			"property1": generate.Schema{
				Type: "string",
			},
			"property2": generate.Schema{
				Type: "string",
			},
			"property3": generate.Schema{
				Type: "number",
			},
			"property4": generate.Schema{
				Type:     "integer",
				Format:   "int32",
				Nullable: true,
			},
		},
	}

	msg := messages.NewMessage("SimpleTestResponse", msgType, "", false, msgProperties, nil, []string{"dataObject"})

	tempPath := os.TempDir()
	tempFileName := "simple_test_response.yaml"

	err := generate.GenerateYamlFile(msg, tempPath, tempFileName)

	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err)
	yaml, _ := msg.ToYaml()
	assert.Equal(t, yaml, string(file))
}

func TestMessage_ToYaml_Complex_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()
	msgType := "object"
	msgProperties["dataObject"] = generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4", "objProp", "arrayProp"},
		Properties: generate.PropertiesMap{
			"property1": generate.Schema{
				Type: "string",
			},
			"property2": generate.Schema{
				Type: "string",
			},
			"property3": generate.Schema{
				Type: "number",
			},
			"property4": generate.Schema{
				Type:     "integer",
				Format:   "int32",
				Nullable: true,
			},
			"objProp": generate.Schema{
				Type: "object",
				Properties: generate.PropertiesMap{
					"objProp1": generate.Schema{
						Type: "string",
					},
					"objProp2": generate.Schema{
						Type: "string",
					},
					"objProp3": generate.Schema{
						Type: "number",
					},
					"objProp4": generate.Schema{
						Type:     "integer",
						Format:   "int32",
						Nullable: true,
					},
				},
				Required: []string{"objProp1", "objProp2", "objProp3", "objProp4"},
			},
			"arrayProp": generate.Schema{
				Type: "array",
				Items: &generate.Schema{
					Type: "object",
					Properties: generate.PropertiesMap{
						"arrayProp1": generate.Schema{
							Type: "string",
						},
						"arrayProp2": generate.Schema{
							Type: "string",
						},
						"arrayProp3": generate.Schema{
							Type: "number",
						},
						"arrayProp4": generate.Schema{
							Type:     "integer",
							Format:   "int32",
							Nullable: true,
						},
					},
					Required: []string{"arrayProp1", "arrayProp2", "arrayProp3", "arrayProp4"},
				},
			},
		},
	}

	msg := messages.NewMessage("ComplexTestResponse", msgType, "", false, msgProperties, nil, []string{"dataObject"})

	yaml, err := msg.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "complex_test_response.yaml"

	err = generate.GenerateYamlFile(msg, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err)
	assert.Equal(t, yaml, string(file))
}

func TestMessage_ToYaml_Empty_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()

	msgType := "object"

	msg := messages.NewMessage("EmptyTestResponse", msgType, "", false, msgProperties, nil, nil)

	yaml, err := msg.ToYaml()
	assert.Nil(t, err)

	tempPath := os.TempDir()
	tempFileName := "empty_test_response.yaml"

	err = generate.GenerateYamlFile(msg, tempPath, tempFileName)
	assert.Nil(t, err)
	assert.FileExists(t, fmt.Sprintf("%s/%s", tempPath, tempFileName))

	// Unmarshal the generated yaml file
	// and compare it with the expected yaml
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, tempFileName))
	assert.Nil(t, err)
	assert.Equal(t, yaml, string(file))
}
