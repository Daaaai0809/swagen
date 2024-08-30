package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen/config"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/models"
)

type ModelCommandParams struct {
	FileName  string
	Title     string
	Type      string
	Properties map[string]generate.Schema
}

func ModelCommandHandler(params ModelCommandParams, dir string) error {
	var path string

	c := config.GetConfig()

	path = fmt.Sprintf("%s/%s", c.GetModelDir(), dir)

	m := models.NewModelSchema(params.Title, params.Type, params.Properties)

	if err := generate.GenerateYamlFile(m, path, params.FileName); err != nil {
		return err
	}

	return nil
}
