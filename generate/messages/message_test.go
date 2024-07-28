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
	msgProperties["dataObject"] = generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4"},
		Properties: map[string]generate.Schema{
			"property1": {
				Type: "string",
			},
			"property2": {
				Type: "string",
			},
			"property3": {
				Type: "number",
			},
			"property4": {
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
	msgProperties["dataObject"] = generate.Schema{
		Type:     "object",
		Required: []string{"property1", "property2", "property3", "property4", "objProp", "arrayProp"},
		Properties: map[string]generate.Schema{
			"property1": {
				Type: "string",
			},
			"property2": {
				Type: "string",
			},
			"property3": {
				Type: "number",
			},
			"property4": {
				Type:     "integer",
				Format:   "int32",
				Nullable: true,
			},
			"objProp": {
				Type: "object",
				Properties: map[string]generate.Schema{
					"objProp1": {
						Type: "string",
					},
					"objProp2": {
						Type: "string",
					},
					"objProp3": {
						Type: "number",
					},
					"objProp4": {
						Type:     "integer",
						Format:   "int32",
						Nullable: true,
					},
				},
				Required: []string{"objProp1", "objProp2", "objProp3", "objProp4"},
			},
			"arrayProp": {
				Type: "array",
				Items: &generate.Schema{
					Type: "object",
					Properties: map[string]generate.Schema{
						"arrayProp1": {
							Type: "string",
						},
						"arrayProp2": {
							Type: "string",
						},
						"arrayProp3": {
							Type: "number",
						},
						"arrayProp4": {
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
