package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen"
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
	Security    []generate.Security
	RequestBody generate.RequestBody
	Parameters  generate.Parameters
	Responses   generate.Responses
}

func PathCommandHandler(params PathCommandParams, dir string) error {
	var path string
	var s generate.ISwaggerSchema = nil

	c := swagen.GetConfig()

	switch params.Method {
	case constant.GET_FILE:
		path = fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewGetPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.Parameters, params.Responses)
	case constant.POST_FILE:
		path = fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewPostPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.RequestBody, params.Parameters, params.Responses)
	case constant.PUT_FILE:
		path = fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewPutPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.RequestBody, params.Parameters, params.Responses)
	case constant.DELETE_FILE:
		path = fmt.Sprintf("%s/%s", c.GetPathDir(), dir)

		s = methods.NewDeletePathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.RequestBody, params.Parameters, params.Responses)
	default:
		return fmt.Errorf("invalid command: %s", params.Method)
	}

	if path == "" {
		return fmt.Errorf("do not have a empty path")
	}

	if err := generate.GenerateYamlFile(s, path, params.FileName); err != nil {
		return err
	}

	return nil
}
