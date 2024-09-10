package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/messages"
)

type MessageCommandParams struct {
	MessageName string
	FileName    string
	Type        string
	Format      string
	Nullable    bool
	Properties  messages.MessageProperties
	Items       *generate.Schema
	Required    []string
}

func MessageCommandHandler(params MessageCommandParams, dir string) error {
	var path string

	c := swagen.GetConfig()

	path = fmt.Sprintf("%s/%s", c.GetMessageDir(), dir)

	m := messages.NewMessage(params.MessageName, params.Type, params.Format, params.Nullable, params.Properties, params.Items, params.Required)

	if err := generate.GenerateYamlFile(m, path, params.FileName); err != nil {
		return err
	}

	return nil
}
