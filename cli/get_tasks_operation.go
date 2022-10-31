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

// makeOperationHandlerGetTasksCmd returns a cmd to handle operation getTasks
func makeOperationHandlerGetTasksCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetTasks",
		Short: `get tasks created by the user`,
		RunE:  runOperationHandlerGetTasks,
	}

	if err := registerOperationHandlerGetTasksParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHandlerGetTasks uses cmd flags to call endpoint api
func runOperationHandlerGetTasks(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := handler.NewGetTasksParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHandlerGetTasksResult(appCli.Handler.GetTasks(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHandlerGetTasksParamFlags registers all flags needed to fill params
func registerOperationHandlerGetTasksParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationHandlerGetTasksResult parses request result and return the string content
func parseOperationHandlerGetTasksResult(resp0 *handler.GetTasksOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getTasksOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*handler.GetTasksBadRequest)
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
		resp2, ok := iResp2.(*handler.GetTasksForbidden)
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
		resp3, ok := iResp3.(*handler.GetTasksInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response getTasksOK is not supported by go-swagger cli yet.

	return "", nil
}
