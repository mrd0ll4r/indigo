// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.temp.transferAccount

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// TempTransferAccount_Input is the input argument to a com.atproto.temp.transferAccount call.
type TempTransferAccount_Input struct {
	Did    string      `json:"did" cborgen:"did"`
	Handle string      `json:"handle" cborgen:"handle"`
	PlcOp  interface{} `json:"plcOp" cborgen:"plcOp"`
}

// TempTransferAccount_Output is the output of a com.atproto.temp.transferAccount call.
type TempTransferAccount_Output struct {
	AccessJwt  string `json:"accessJwt" cborgen:"accessJwt"`
	Did        string `json:"did" cborgen:"did"`
	Handle     string `json:"handle" cborgen:"handle"`
	RefreshJwt string `json:"refreshJwt" cborgen:"refreshJwt"`
}

// TempTransferAccount calls the XRPC method "com.atproto.temp.transferAccount".
func TempTransferAccount(ctx context.Context, c *xrpc.Client, input *TempTransferAccount_Input) (*TempTransferAccount_Output, error) {
	var out TempTransferAccount_Output
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.temp.transferAccount", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
