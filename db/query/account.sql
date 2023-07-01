-- name: CreateAccount :one
insert into "accounts" (account_holder_id, account_number)
values ($1, $2)
returning id;

-- name: GetAccountByAccountNumber :one
select *
from "accounts"
where account_number = $1;

-- name: GetAccountWithHolder :one
select "accounts".id,
       "accounts".account_number,
       "accounts".balance,
       "accounts".created_at,
       "account_holder".id as "holder_id",
       "account_holder".first_name,
       "account_holder".last_name,
       "account_holder".email,
       "account_holder".phone,
       "account_holder".address
from "accounts"
         join "account_holders" as "account_holder"
              on "accounts".account_holder_id = "account_holder".id
where "accounts".account_number = $1;

-- name: DoesAccountNumberExist :one
select exists(select 1 from "accounts" where account_number = $1);

-- name: GetAccountsFromHolder :many
select *
from "accounts"
where account_holder_id = $1
limit $2 offset $3;