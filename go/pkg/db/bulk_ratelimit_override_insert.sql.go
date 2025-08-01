// Code generated by sqlc bulk insert plugin. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"strings"
)

// bulkInsertRatelimitOverride is the base query for bulk insert
const bulkInsertRatelimitOverride = `INSERT INTO ratelimit_overrides ( id, workspace_id, namespace_id, identifier, ` + "`" + `limit` + "`" + `, duration, async, created_at_m ) VALUES %s ON DUPLICATE KEY UPDATE
    ` + "`" + `limit` + "`" + ` = VALUES(` + "`" + `limit` + "`" + `),
    duration = VALUES(duration),
    async = VALUES(async),
    updated_at_m = ?`

// InsertRatelimitOverrides performs bulk insert in a single query
func (q *BulkQueries) InsertRatelimitOverrides(ctx context.Context, db DBTX, args []InsertRatelimitOverrideParams) error {

	if len(args) == 0 {
		return nil
	}

	// Build the bulk insert query
	valueClauses := make([]string, len(args))
	for i := range args {
		valueClauses[i] = "( ?, ?, ?, ?, ?, ?, false, ? )"
	}

	bulkQuery := fmt.Sprintf(bulkInsertRatelimitOverride, strings.Join(valueClauses, ", "))

	// Collect all arguments
	var allArgs []any
	for _, arg := range args {
		allArgs = append(allArgs, arg.ID)
		allArgs = append(allArgs, arg.WorkspaceID)
		allArgs = append(allArgs, arg.NamespaceID)
		allArgs = append(allArgs, arg.Identifier)
		allArgs = append(allArgs, arg.Limit)
		allArgs = append(allArgs, arg.Duration)
		allArgs = append(allArgs, arg.CreatedAt)
	}

	// Add ON DUPLICATE KEY UPDATE parameters (only once, not per row)
	if len(args) > 0 {
		allArgs = append(allArgs, args[0].UpdatedAt)
	}

	// Execute the bulk insert
	_, err := db.ExecContext(ctx, bulkQuery, allArgs...)
	return err
}
