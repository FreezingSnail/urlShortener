// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createUrl = `-- name: CreateUrl :one
INSERT INTO urls (
  url, shorturl, userid, createdate
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, url, shorturl, userid, createdate
`

type CreateUrlParams struct {
	Url        string
	Shorturl   string
	Userid     sql.NullInt64
	Createdate sql.NullInt64
}

func (q *Queries) CreateUrl(ctx context.Context, arg CreateUrlParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createUrl,
		arg.Url,
		arg.Shorturl,
		arg.Userid,
		arg.Createdate,
	)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Shorturl,
		&i.Userid,
		&i.Createdate,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  ?, ?
)
RETURNING id, name, email
`

type CreateUserParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteUrl = `-- name: DeleteUrl :exec
DELETE FROM urls
WHERE id = ?
`

func (q *Queries) DeleteUrl(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUrl, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getURL = `-- name: GetURL :one
SELECT id, url, shorturl, userid, createdate FROM urls
WHERE url = ? LIMIT 1
`

func (q *Queries) GetURL(ctx context.Context, url string) (Url, error) {
	row := q.db.QueryRowContext(ctx, getURL, url)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Shorturl,
		&i.Userid,
		&i.Createdate,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const listUrls = `-- name: ListUrls :many
SELECT id, url, shorturl, userid, createdate FROM urls
ORDER BY url
`

func (q *Queries) ListUrls(ctx context.Context) ([]Url, error) {
	rows, err := q.db.QueryContext(ctx, listUrls)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Url
	for rows.Next() {
		var i Url
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Shorturl,
			&i.Userid,
			&i.Createdate,
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

const listUsers = `-- name: ListUsers :many
SELECT id, name, email FROM users
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
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