// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_override_list_by_namespace_id.sql

package db

import (
	"context"
)

const listRatelimitOverridesByNamespaceID = `-- name: ListRatelimitOverridesByNamespaceID :many
SELECT id, workspace_id, namespace_id, identifier, ` + "`" + `limit` + "`" + `, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
WHERE
workspace_id = ?
AND namespace_id = ?
AND deleted_at_m IS NULL
AND id >= ?
ORDER BY id ASC
LIMIT ?
`

type ListRatelimitOverridesByNamespaceIDParams struct {
	WorkspaceID string `db:"workspace_id"`
	NamespaceID string `db:"namespace_id"`
	CursorID    string `db:"cursor_id"`
	Limit       int32  `db:"limit"`
}

// ListRatelimitOverridesByNamespaceID
//
//	SELECT id, workspace_id, namespace_id, identifier, `limit`, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
//	WHERE
//	workspace_id = ?
//	AND namespace_id = ?
//	AND deleted_at_m IS NULL
//	AND id >= ?
//	ORDER BY id ASC
//	LIMIT ?
func (q *Queries) ListRatelimitOverridesByNamespaceID(ctx context.Context, db DBTX, arg ListRatelimitOverridesByNamespaceIDParams) ([]RatelimitOverride, error) {
	rows, err := db.QueryContext(ctx, listRatelimitOverridesByNamespaceID,
		arg.WorkspaceID,
		arg.NamespaceID,
		arg.CursorID,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RatelimitOverride
	for rows.Next() {
		var i RatelimitOverride
		if err := rows.Scan(
			&i.ID,
			&i.WorkspaceID,
			&i.NamespaceID,
			&i.Identifier,
			&i.Limit,
			&i.Duration,
			&i.Async,
			&i.Sharding,
			&i.CreatedAtM,
			&i.UpdatedAtM,
			&i.DeletedAtM,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
