-- name: CreateAccount :one
insert into "accounts" (account_holder_id, account_number)
values ($1, $2)
returning id;

-- name: GetAccountByAccountNumber :one
select *
from "accounts"
where account_number = $1;

-- name: GetAccountWithHolder :one
select *
from "accounts"
         join "account_holders" as "account_holder"
              on "accounts".account_holder_id = "account_holder".id
where "accounts".account_number = $1;