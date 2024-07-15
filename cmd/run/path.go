package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen/config"
	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/methods"
)

type PathCommandParams struct {
	Method      string
	FileName    string
	OperationID string
	Summary     string
	Description string
	Tags        []string
	Security    []methods.Security
	RequestBody methods.RequestBody
	Parameters  methods.Parameters
	Responses   methods.Responses
}

func PathCommandHandler(params PathCommandParams, dir string) error {
	var s generate.IPathSchema = nil

	c := config.GetConfig()

	switch params.Method {
	case constant.GET_FILE:
		path := fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewGetPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.Parameters, params.Responses)
		if err := generate.GeneratePathYamlFile(s, path, params.FileName); err != nil {
			return err
		}
		return nil
	case constant.POST_FILE:
		path := fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewPostPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.RequestBody, params.Parameters, params.Responses)
		if err := generate.GeneratePathYamlFile(s, path, params.FileName); err != nil {
			return err
		}
		return nil
	case constant.PUT_FILE:
	case constant.DELETE_FILE:
	default:
		return fmt.Errorf("invalid command: %s", params.Method)
	}

	return nil
}
