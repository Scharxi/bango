create table if not exists "account_holders"
(
    id         serial primary key not null,
    first_name varchar(100)       not null,
    last_name  varchar(100)       not null,
    email      varchar(100)       not null,
    phone      varchar(100)       not null,
    address    varchar(100)       not null,
    created_at timestamp          not null default now()
);


create table if not exists "accounts"
(
    id                serial primary key not null,
    account_holder_id serial             not null,
    account_number    serial             not null,
    balance           decimal            not null default 0.00,
    created_at        timestamp          not null default now(),
    FOREIGN KEY (account_holder_id) REFERENCES account_holders (id) on delete cascade
);
