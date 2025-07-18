// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: key_find_credits.sql

package db

import (
	"context"
	"database/sql"
)

const findKeyCredits = `-- name: FindKeyCredits :one
SELECT remaining_requests FROM ` + "`" + `keys` + "`" + ` k WHERE k.id = ?
`

// FindKeyCredits
//
//	SELECT remaining_requests FROM `keys` k WHERE k.id = ?
func (q *Queries) FindKeyCredits(ctx context.Context, db DBTX, id string) (sql.NullInt32, error) {
	row := db.QueryRowContext(ctx, findKeyCredits, id)
	var remaining_requests sql.NullInt32
	err := row.Scan(&remaining_requests)
	return remaining_requests, err
}
