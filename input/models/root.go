package models

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type IModelInput interface {
	ReadAll()
}

type ModelInputs struct {
	Cmd     *cobra.Command
	Title   string
	Type    string
	Properties map[string]generate.Schema
}

func NewModelInputs(cmd *cobra.Command) *ModelInputs {
	return &ModelInputs{
		Cmd: cmd,
	}
}

// The ReadAll method reads all the input required to define a schema.
func (m *ModelInputs) ReadAll() {
	m.ReadTitle()
	m.ReadType()
	m.ReadProperties()
}

func (m *ModelInputs) SetTitle(title string) {
	m.Title = title
}

func (m *ModelInputs) GetTitle() string {
	return m.Title
}

func (m *ModelInputs) ReadTitle() {
	scanner := bufio.NewScanner(os.Stdin)

	println("Enter the title of the model:")

	scanner.Scan()

	m.SetTitle(scanner.Text())
}

func (m *ModelInputs) SetType(t string) {
	m.Type = t
}

func (m *ModelInputs) GetType() string {
	return m.Type
}

func (m *ModelInputs) ReadType() {
	t, err := input.SingleSelect("Select the type of the model:", constant.SchemaTypeList)
	if err != nil {
		println("Error reading type: ", err)
		os.Exit(1)
	}

	m.SetType(t)
}

func (m *ModelInputs) SetProperties(properties map[string]generate.Schema) {
	m.Properties = properties
}

func (m *ModelInputs) GetProperties() map[string]generate.Schema {
	return m.Properties
}

func (m *ModelInputs) ReadProperties() {
	properties := make(map[string]generate.Schema)

	for {
		schema := generate.Schema{}
		
		println("Enter the field name: ")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		fieldName := scanner.Text()

		schema.ReadType(m.Cmd)

		switch schema.Type {
		case constant.STRING_TYPE, constant.NUMBER_TYPE, constant.INTEGER_TYPE:
			schema.ReadFormat(m.Cmd, schema.Type)
		case constant.OBJECT_TYPE:
			schema.ReadProperties(m.Cmd, scanner, true)
		case constant.ARRAY_TYPE:
			schema.ReadItems(m.Cmd, scanner, true)
		}

		if ok := input.YesNoPrompt(m.Cmd, "Is the field nullable?"); ok {
			schema.Nullable = true
		}

		properties[fieldName] = schema

		if ok := input.YesNoPrompt(m.Cmd, "Do you want to add more fields?"); !ok {
			break
		}
	}

	m.SetProperties(properties)
}
