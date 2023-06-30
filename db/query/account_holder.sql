-- name: CreateAccountHolder :one
-- description: Create an account holder
insert into "account_holders" (first_name, last_name, email, phone, address) values ($1, $2, $3, $4, $5) returning id;