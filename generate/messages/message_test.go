package messages_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/messages"
)

const SIMPLE_TEST_RESPONSE_FILE_PATH = "tests/generate/messages/SimpleTestResponse.yaml"
const COMPLEX_TEST_RESPONSE_FILE_PATH = "tests/generate/messages/ComplexTestResponse.yaml"
const EMPTY_TEST_RESPONSE_FILE_PATH = "tests/generate/messages/EmptyTestResponse.yaml"

var (
	expectedYaml_Simple_Response  string
	expectedYaml_Complex_Response string
	expectedYaml_Empty_Response   string
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

	rootDir := currentDir[:len(currentDir)-len("generate/messages")]

	// NOTE: get TestFiles from root/tests/generate/messages
	expectedYaml_Simple_Response_Byte, err := os.ReadFile(fmt.Sprint(rootDir, SIMPLE_TEST_RESPONSE_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Complex_Response_Byte, err := os.ReadFile(fmt.Sprint(rootDir, COMPLEX_TEST_RESPONSE_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Empty_Response_Byte, err := os.ReadFile(fmt.Sprint(rootDir, EMPTY_TEST_RESPONSE_FILE_PATH))
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	expectedYaml_Simple_Response = string(expectedYaml_Simple_Response_Byte)
	expectedYaml_Complex_Response = string(expectedYaml_Complex_Response_Byte)
	expectedYaml_Empty_Response = string(expectedYaml_Empty_Response_Byte)

	os.Exit(m.Run())
}

func TestMessage_ToYaml_Simple_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()

	msgType := "object"
	msgProperties["dataObject"] = &generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4"},
		Properties: generate.PropertiesMap{
			"property1": &generate.Schema{
				Type: "string",
			},
			"property2": &generate.Schema{
				Type: "string",
			},
			"property3": &generate.Schema{
				Type: "number",
			},
			"property4": &generate.Schema{
				Type:     "integer",
				Format:   "int32",
				Nullable: true,
			},
		},
	}

	msg := messages.NewMessage("SimpleTestResponse", msgType, "", false, msgProperties, nil, []string{"dataObject"})

	yaml, err := msg.ToYaml()

	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Simple_Response, yaml)
}

func TestMessage_ToYaml_Complex_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()
	msgType := "object"
	msgProperties["dataObject"] = &generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4", "objProp", "arrayProp"},
		Properties: generate.PropertiesMap{
			"property1": &generate.Schema{
				Type: "string",
			},
			"property2": &generate.Schema{
				Type: "string",
			},
			"property3": &generate.Schema{
				Type: "number",
			},
			"property4": &generate.Schema{
				Type:     "integer",
				Format:   "int32",
				Nullable: true,
			},
			"objProp": &generate.Schema{
				Type: "object",
				Properties: generate.PropertiesMap{
					"objProp1": &generate.Schema{
						Type: "string",
					},
					"objProp2": &generate.Schema{
						Type: "string",
					},
					"objProp3": &generate.Schema{
						Type: "number",
					},
					"objProp4": &generate.Schema{
						Type:     "integer",
						Format:   "int32",
						Nullable: true,
					},
				},
				Required: []string{"objProp1", "objProp2", "objProp3", "objProp4"},
			},
			"arrayProp": &generate.Schema{
				Type: "array",
				Items: &generate.Schema{
					Type: "object",
					Properties: generate.PropertiesMap{
						"arrayProp1": &generate.Schema{
							Type: "string",
						},
						"arrayProp2": &generate.Schema{
							Type: "string",
						},
						"arrayProp3": &generate.Schema{
							Type: "number",
						},
						"arrayProp4": &generate.Schema{
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
	assert.Equal(t, expectedYaml_Complex_Response, yaml)
}

func TestMessage_ToYaml_Empty_Response(t *testing.T) {
	msgProperties := messages.NewMessageProperties()

	msgType := "object"

	msg := messages.NewMessage("EmptyTestResponse", msgType, "", false, msgProperties, nil, nil)

	yaml, err := msg.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, expectedYaml_Empty_Response, yaml)
}
