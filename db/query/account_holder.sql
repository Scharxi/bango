-- name: CreateAccountHolder :one
-- description: Create an account holder
insert into "account_holders" (first_name, last_name, email, phone, address)
values ($1, $2, $3, $4, $5)
returning id;

-- name: DoesEmailExist :one
-- description: Check if email exists
select exists(select 1 from "account_holders" where email = $1);

-- name: DoesPhoneNumberExist :one
-- description: Check if phone number exists
select exists(select 1 from "account_holders" where phone = $1);