package messages

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/messages"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type IMessageInput interface {
	ReadAll()
}

type MessageInputs struct {
	Cmd         *cobra.Command
	MessageName string
	Type        string
	Format      string
	Nullable    bool
	Properties  messages.MessageProperties
	Items       *generate.Schema
	Required    []string
}

func NewMessageInputs(cmd *cobra.Command) *MessageInputs {
	return &MessageInputs{
		Cmd: cmd,
	}
}

// The ReadAll method reads all the input required to define a schema.
func (m *MessageInputs) ReadAll() {
	m.ReadMessageName()
	m.ReadType()
	if m.Type != constant.OBJECT_TYPE && m.Type != constant.ARRAY_TYPE {
		m.ReadFormat()
	}
	m.ReadNullable()

	if m.Type == constant.OBJECT_TYPE {
		m.ReadMessageProperties()
	}

	if m.Type == constant.ARRAY_TYPE {
		m.ReadItems()
	}
}

func (m *MessageInputs) SetMessageName(messageName string) {
	m.MessageName = messageName
}

func (m *MessageInputs) GetMessageName() string {
	return m.MessageName
}

// The ReadMessageName method takes input from the CLI to define the message name for the schema.
func (m *MessageInputs) ReadMessageName() {
	scanner := bufio.NewScanner(os.Stdin)

	println("Enter the message name: ")

	scanner.Scan()

	m.SetMessageName(scanner.Text())
}

func (m *MessageInputs) SetType(type_ string) {
	m.Type = type_
}

func (m *MessageInputs) GetType() string {
	return m.Type
}

func (m *MessageInputs) ReadType() {
	t, err := input.SingleSelect("Select the type of the message", constant.SchemaTypeList)
	if err != nil {
		println("Error reading type: ", err)
		os.Exit(1)
	}

	m.SetType(t)
}

func (m *MessageInputs) SetFormat(format string) {
	m.Format = format
}

func (m *MessageInputs) GetFormat() string {
	return m.Format
}

func (m *MessageInputs) ReadFormat() {
	switch m.Type {
	case constant.STRING_TYPE:
		f, err := input.SingleSelect("Select the format of the message", constant.FormatStringList)
		if err != nil {
			println("Error reading format: ", err)
			os.Exit(1)
		}

		m.SetFormat(f)
	case constant.NUMBER_TYPE:
		f, err := input.SingleSelect("Select the format of the message", constant.FormatNumberList)
		if err != nil {
			println("Error reading format: ", err)
			os.Exit(1)
		}

		m.SetFormat(f)
	case constant.INTEGER_TYPE:
		f, err := input.SingleSelect("Select the format of the message", constant.FormatIntegerList)
		if err != nil {
			println("Error reading format: ", err)
			os.Exit(1)
		}

		m.SetFormat(f)
	}
}

func (m *MessageInputs) SetNullable(nullable bool) {
	m.Nullable = nullable
}

func (m *MessageInputs) GetNullable() bool {
	return m.Nullable
}

func (m *MessageInputs) ReadNullable() {
	if ok := input.YesNoPrompt(m.Cmd, "Is the message nullable?"); ok {
		m.SetNullable(true)
	}
}

func (m *MessageInputs) SetRequired(required []string) {
	m.Required = required
}

func (m *MessageInputs) GetRequired() []string {
	return m.Required
}

func (m *MessageInputs) SetMessageProperties(messageProperties messages.MessageProperties) {
	m.Properties = messageProperties
}

func (m *MessageInputs) GetMessageProperties() messages.MessageProperties {
	return m.Properties
}

// The ReadMessageProperties method takes input from the CLI to define the message root for the schema.
func (m *MessageInputs) ReadMessageProperties() {
	msgRoot := messages.NewMessageProperties()

	for {
		schema := generate.NewSchema()

		println("Enter the field name: ")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		fieldName := scanner.Text()

		schema.ReadType(m.Cmd)

		switch schema.Type {
		case constant.STRING_TYPE, constant.NUMBER_TYPE, constant.INTEGER_TYPE:
			schema.ReadFormat(m.Cmd, schema.Type)
		case constant.OBJECT_TYPE:
			schema.ReadProperties(m.Cmd, scanner, false)
		case constant.ARRAY_TYPE:
			schema.ReadItems(m.Cmd, scanner, false)
		}

		if ok := input.YesNoPrompt(m.Cmd, "Is the field required?"); ok {
			m.Required = append(m.Required, fieldName)
		}

		if ok := input.YesNoPrompt(m.Cmd, "Is the field nullable?"); ok {
			schema.Nullable = true
		}

		msgRoot[fieldName] = *schema

		if ok := input.YesNoPrompt(m.Cmd, "Do you want to add more fields?"); !ok {
			break
		}
	}

	m.SetMessageProperties(msgRoot)
}

func (m *MessageInputs) SetItems(items *generate.Schema) {
	m.Items = items
}

func (m *MessageInputs) GetItems() *generate.Schema {
	return m.Items
}

func (m *MessageInputs) ReadItems() {
	schema := generate.NewSchema()

	schema.ReadType(m.Cmd)

	scanner := bufio.NewScanner(os.Stdin)

	switch schema.Type {
	case constant.STRING_TYPE, constant.NUMBER_TYPE, constant.INTEGER_TYPE:
		schema.ReadFormat(m.Cmd, schema.Type)
	case constant.OBJECT_TYPE:
		schema.ReadProperties(m.Cmd, scanner, false)
	case constant.ARRAY_TYPE:
		schema.ReadItems(m.Cmd, scanner, false)
	}

	m.SetItems(schema)
}
