// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: key_permission_delete_all_by_key_id.sql

package db

import (
	"context"
)

const deleteAllKeyPermissionsByKeyID = `-- name: DeleteAllKeyPermissionsByKeyID :exec
DELETE FROM keys_permissions
WHERE key_id = ?
`

// DeleteAllKeyPermissionsByKeyID
//
//	DELETE FROM keys_permissions
//	WHERE key_id = ?
func (q *Queries) DeleteAllKeyPermissionsByKeyID(ctx context.Context, db DBTX, keyID string) error {
	_, err := db.ExecContext(ctx, deleteAllKeyPermissionsByKeyID, keyID)
	return err
}
