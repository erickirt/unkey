// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: key_encryption_insert.sql

package db

import (
	"context"
)

const insertKeyEncryption = `-- name: InsertKeyEncryption :exec
INSERT INTO encrypted_keys
(workspace_id, key_id, encrypted, encryption_key_id, created_at)
VALUES (?, ?, ?, ?, ?)
`

type InsertKeyEncryptionParams struct {
	WorkspaceID     string `db:"workspace_id"`
	KeyID           string `db:"key_id"`
	Encrypted       string `db:"encrypted"`
	EncryptionKeyID string `db:"encryption_key_id"`
	CreatedAt       int64  `db:"created_at"`
}

// InsertKeyEncryption
//
//	INSERT INTO encrypted_keys
//	(workspace_id, key_id, encrypted, encryption_key_id, created_at)
//	VALUES (?, ?, ?, ?, ?)
func (q *Queries) InsertKeyEncryption(ctx context.Context, db DBTX, arg InsertKeyEncryptionParams) error {
	_, err := db.ExecContext(ctx, insertKeyEncryption,
		arg.WorkspaceID,
		arg.KeyID,
		arg.Encrypted,
		arg.EncryptionKeyID,
		arg.CreatedAt,
	)
	return err
}
