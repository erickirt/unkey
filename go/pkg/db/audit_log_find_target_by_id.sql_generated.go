// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: audit_log_find_target_by_id.sql

package db

import (
	"context"
)

const findAuditLogTargetByID = `-- name: FindAuditLogTargetByID :many
SELECT audit_log_target.workspace_id, audit_log_target.bucket_id, audit_log_target.bucket, audit_log_target.audit_log_id, audit_log_target.display_name, audit_log_target.type, audit_log_target.id, audit_log_target.name, audit_log_target.meta, audit_log_target.created_at, audit_log_target.updated_at, audit_log.id, audit_log.workspace_id, audit_log.bucket, audit_log.bucket_id, audit_log.event, audit_log.time, audit_log.display, audit_log.remote_ip, audit_log.user_agent, audit_log.actor_type, audit_log.actor_id, audit_log.actor_name, audit_log.actor_meta, audit_log.created_at, audit_log.updated_at
FROM audit_log_target
JOIN audit_log ON audit_log.id = audit_log_target.audit_log_id
WHERE audit_log_target.id = ?
`

type FindAuditLogTargetByIDRow struct {
	AuditLogTarget AuditLogTarget `db:"audit_log_target"`
	AuditLog       AuditLog       `db:"audit_log"`
}

// FindAuditLogTargetByID
//
//	SELECT audit_log_target.workspace_id, audit_log_target.bucket_id, audit_log_target.bucket, audit_log_target.audit_log_id, audit_log_target.display_name, audit_log_target.type, audit_log_target.id, audit_log_target.name, audit_log_target.meta, audit_log_target.created_at, audit_log_target.updated_at, audit_log.id, audit_log.workspace_id, audit_log.bucket, audit_log.bucket_id, audit_log.event, audit_log.time, audit_log.display, audit_log.remote_ip, audit_log.user_agent, audit_log.actor_type, audit_log.actor_id, audit_log.actor_name, audit_log.actor_meta, audit_log.created_at, audit_log.updated_at
//	FROM audit_log_target
//	JOIN audit_log ON audit_log.id = audit_log_target.audit_log_id
//	WHERE audit_log_target.id = ?
func (q *Queries) FindAuditLogTargetByID(ctx context.Context, db DBTX, id string) ([]FindAuditLogTargetByIDRow, error) {
	rows, err := db.QueryContext(ctx, findAuditLogTargetByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAuditLogTargetByIDRow
	for rows.Next() {
		var i FindAuditLogTargetByIDRow
		if err := rows.Scan(
			&i.AuditLogTarget.WorkspaceID,
			&i.AuditLogTarget.BucketID,
			&i.AuditLogTarget.Bucket,
			&i.AuditLogTarget.AuditLogID,
			&i.AuditLogTarget.DisplayName,
			&i.AuditLogTarget.Type,
			&i.AuditLogTarget.ID,
			&i.AuditLogTarget.Name,
			&i.AuditLogTarget.Meta,
			&i.AuditLogTarget.CreatedAt,
			&i.AuditLogTarget.UpdatedAt,
			&i.AuditLog.ID,
			&i.AuditLog.WorkspaceID,
			&i.AuditLog.Bucket,
			&i.AuditLog.BucketID,
			&i.AuditLog.Event,
			&i.AuditLog.Time,
			&i.AuditLog.Display,
			&i.AuditLog.RemoteIp,
			&i.AuditLog.UserAgent,
			&i.AuditLog.ActorType,
			&i.AuditLog.ActorID,
			&i.AuditLog.ActorName,
			&i.AuditLog.ActorMeta,
			&i.AuditLog.CreatedAt,
			&i.AuditLog.UpdatedAt,
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
