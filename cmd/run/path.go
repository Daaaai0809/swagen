package run

import (
	"fmt"

	"github.com/Daaaai0809/swagen/config"
	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/Daaaai0809/swagen/generate/methods"
	"github.com/Daaaai0809/swagen/generate/types"
)

type PathCommandParams struct {
	Method      string
	FileName    string
	OperationID string
	Summary     string
	Description string
	Tags        []string
	Security    []types.Security
	Parameters  types.Parameters
	Responses   types.Responses
}

func PathCommandHandler(params PathCommandParams) error {
	var s generate.IPathSchema = nil

	c := config.GetConfig()

	switch params.Method {
	case constant.GET_FILE:
		s = methods.NewGetPathSchema(params.OperationID, params.Summary, params.Description, params.Tags, params.Security, params.Parameters, params.Responses)
		if err := generate.GeneratePathYamlFile(s, c.GetPathDir(), params.FileName); err != nil {
			return err
		}
		return nil
	case constant.POST_FILE:
	case constant.PUT_FILE:
	case constant.DELETE_FILE:
	default:
		return fmt.Errorf("invalid command: %s", params.Method)
	}

	return nil
}
