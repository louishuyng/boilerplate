// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: create_example.sql

package repository

import (
	"context"
)

const createExample = `-- name: CreateExample :one
INSERT INTO example (id, name)
VALUES (
    gen_random_uuid(),
    $1
)
RETURNING id, name
`

func (q *Queries) CreateExample(ctx context.Context, name string) (Example, error) {
	row := q.db.QueryRowContext(ctx, createExample, name)
	var i Example
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
