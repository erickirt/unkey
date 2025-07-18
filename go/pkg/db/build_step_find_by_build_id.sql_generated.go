// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: build_step_find_by_build_id.sql

package db

import (
	"context"
)

const findVersionStepsByVersionId = `-- name: FindVersionStepsByVersionId :many
SELECT 
    version_id,
    status,
    message,
    error_message,
    created_at
FROM version_steps 
WHERE version_id = ?
ORDER BY created_at ASC
`

// FindVersionStepsByVersionId
//
//	SELECT
//	    version_id,
//	    status,
//	    message,
//	    error_message,
//	    created_at
//	FROM version_steps
//	WHERE version_id = ?
//	ORDER BY created_at ASC
func (q *Queries) FindVersionStepsByVersionId(ctx context.Context, db DBTX, versionID string) ([]VersionStep, error) {
	rows, err := db.QueryContext(ctx, findVersionStepsByVersionId, versionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VersionStep
	for rows.Next() {
		var i VersionStep
		if err := rows.Scan(
			&i.VersionID,
			&i.Status,
			&i.Message,
			&i.ErrorMessage,
			&i.CreatedAt,
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
