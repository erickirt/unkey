// Code generated by sqlc bulk insert plugin. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"strings"
)

// bulkInsertKeyEncryption is the base query for bulk insert
const bulkInsertKeyEncryption = `INSERT INTO encrypted_keys (workspace_id, key_id, encrypted, encryption_key_id, created_at) VALUES %s`

// InsertKeyEncryptions performs bulk insert in a single query
func (q *BulkQueries) InsertKeyEncryptions(ctx context.Context, db DBTX, args []InsertKeyEncryptionParams) error {

	if len(args) == 0 {
		return nil
	}

	// Build the bulk insert query
	valueClauses := make([]string, len(args))
	for i := range args {
		valueClauses[i] = "(?, ?, ?, ?, ?)"
	}

	bulkQuery := fmt.Sprintf(bulkInsertKeyEncryption, strings.Join(valueClauses, ", "))

	// Collect all arguments
	var allArgs []any
	for _, arg := range args {
		allArgs = append(allArgs, arg.WorkspaceID)
		allArgs = append(allArgs, arg.KeyID)
		allArgs = append(allArgs, arg.Encrypted)
		allArgs = append(allArgs, arg.EncryptionKeyID)
		allArgs = append(allArgs, arg.CreatedAt)
	}

	// Execute the bulk insert
	_, err := db.ExecContext(ctx, bulkQuery, allArgs...)
	return err
}
