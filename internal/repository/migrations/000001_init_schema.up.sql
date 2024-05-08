create table if not exists public.users
(
    id               uuid not null
        primary key,
    first_name       text,
    last_name        text,
    hashed_password  text,
    current_position text,
    position_id      uuid
        constraint position_id
            references public.positions
);

alter table public.employees
    owner to postgres;



create table if not exists public.positions
(
    id            uuid not null
        primary key,
    position_name text,
    employees     text[],
    shifts        text[],
    start_time    varchar(10),
    end_time      varchar(10),
    created_at    timestamp,
    updated_at    timestamp
);

alter table public.positions
    owner to postgres;

