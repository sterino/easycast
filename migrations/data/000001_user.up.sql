begin;

create table user
(
    id            uuid        not null primary key,
    phone_number varchar not null primary key,
    avatar varchar default null,
    phone_verified boolean default false,
    password varchar not null,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

commit;