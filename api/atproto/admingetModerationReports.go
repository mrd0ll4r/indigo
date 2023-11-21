// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.getModerationReports

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminGetModerationReports_Output is the output of a com.atproto.admin.getModerationReports call.
type AdminGetModerationReports_Output struct {
	Cursor  *string                 `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Reports []*AdminDefs_ReportView `json:"reports" cborgen:"reports"`
}

// AdminGetModerationReports calls the XRPC method "com.atproto.admin.getModerationReports".
//
// actionedBy: Get all reports that were actioned by a specific moderator.
// reporters: Filter reports made by one or more DIDs.
// reverse: Reverse the order of the returned records. When true, returns reports in chronological order.
func AdminGetModerationReports(ctx context.Context, c *xrpc.Client, actionType string, actionedBy string, cursor string, ignoreSubjects []string, limit int64, reporters []string, resolved bool, reverse bool, subject string) (*AdminGetModerationReports_Output, error) {
	var out AdminGetModerationReports_Output

	params := map[string]interface{}{
		"actionType":     actionType,
		"actionedBy":     actionedBy,
		"cursor":         cursor,
		"ignoreSubjects": ignoreSubjects,
		"limit":          limit,
		"reporters":      reporters,
		"resolved":       resolved,
		"reverse":        reverse,
		"subject":        subject,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.getModerationReports", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
