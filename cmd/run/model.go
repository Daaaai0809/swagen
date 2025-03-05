package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen/config"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/models"
)

type ModelCommandParams struct {
	ModelName  string
	FileName   string
	Title      string
	Type       string
	Properties generate.IPropertiesMap
}

func ModelCommandHandler(params ModelCommandParams, dir string) error {
	var path string

	c := config.GetConfig()

	path = fmt.Sprintf("%s/%s", c.GetModelDir(), dir)

	props, ok := params.Properties.(generate.PropertiesMap)
	if !ok {
		return fmt.Errorf("failed to convert properties to PropertiesMap")
	}

	m := models.NewModelSchema(params.ModelName, params.Title, params.Type, props)

	if err := generate.GenerateYamlFile(m, path, params.FileName); err != nil {
		return err
	}

	return nil
}
