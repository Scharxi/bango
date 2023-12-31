// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: account_holder.sql

package db

import (
	"context"
)

const createAccountHolder = `-- name: CreateAccountHolder :one
insert into "account_holders" (first_name, last_name, email, phone, address)
values ($1, $2, $3, $4, $5)
returning id
`

type CreateAccountHolderParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

// description: Create an account holder
func (q *Queries) CreateAccountHolder(ctx context.Context, arg CreateAccountHolderParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createAccountHolder,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.Address,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const doesAccountHolderExist = `-- name: DoesAccountHolderExist :one
select exists(select 1 from "account_holders" where id = $1)
`

// description: Check if account holder exists
func (q *Queries) DoesAccountHolderExist(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, doesAccountHolderExist, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const doesEmailExist = `-- name: DoesEmailExist :one
select exists(select 1 from "account_holders" where email = $1)
`

// description: Check if email exists
func (q *Queries) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, doesEmailExist, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const doesPhoneNumberExist = `-- name: DoesPhoneNumberExist :one
select exists(select 1 from "account_holders" where phone = $1)
`

// description: Check if phone number exists
func (q *Queries) DoesPhoneNumberExist(ctx context.Context, phone string) (bool, error) {
	row := q.db.QueryRowContext(ctx, doesPhoneNumberExist, phone)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
