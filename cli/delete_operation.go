// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"gitlab.com/golang-hse-2022/team1/tasks/client/handler"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHandlerDeleteCmd returns a cmd to handle operation delete
func makeOperationHandlerDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "Delete",
		Short: `Delete task by taskId in case it has been assigned to the user`,
		RunE:  runOperationHandlerDelete,
	}

	if err := registerOperationHandlerDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHandlerDelete uses cmd flags to call endpoint api
func runOperationHandlerDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := handler.NewDeleteParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHandlerDeleteResult(appCli.Handler.Delete(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHandlerDeleteParamFlags registers all flags needed to fill params
func registerOperationHandlerDeleteParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationHandlerDeleteResult parses request result and return the string content
func parseOperationHandlerDeleteResult(resp0 *handler.DeleteOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*handler.DeleteBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*handler.DeleteForbidden)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*handler.DeleteNotFound)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*handler.DeleteInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteOK is not supported by go-swagger cli yet.

	return "", nil
}
