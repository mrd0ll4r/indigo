// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.communication.deleteTemplate

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// CommunicationDeleteTemplate_Input is the input argument to a tools.ozone.communication.deleteTemplate call.
type CommunicationDeleteTemplate_Input struct {
	Id string `json:"id" cborgen:"id"`
}

// CommunicationDeleteTemplate calls the XRPC method "tools.ozone.communication.deleteTemplate".
func CommunicationDeleteTemplate(ctx context.Context, c *xrpc.Client, input *CommunicationDeleteTemplate_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "tools.ozone.communication.deleteTemplate", nil, input, nil); err != nil {
		return err
	}

	return nil
}