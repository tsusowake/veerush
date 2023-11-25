// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package generated

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users default
values
returning id, created_at
`

func (q *Queries) CreateUser(ctx context.Context) (User, error) {
	row := q.db.QueryRow(ctx, createUser)
	var i User
	err := row.Scan(&i.ID, &i.CreatedAt)
	return i, err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
delete
from users
where id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUserByID, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
select id, created_at
from users
where id = $1
limit 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.CreatedAt)
	return i, err
}
