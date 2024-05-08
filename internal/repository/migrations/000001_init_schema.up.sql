create table if not exists public.positions
(
    id            bigserial
        primary key,
    position_name text,
    created_at    timestamp,
    updated_at    timestamp
);

alter table public.positions
    owner to postgres;

create table if not exists public.users
(
    id               bigserial
        primary key,
    first_name       text,
    last_name        text,
    email            text,
    password         text,
    current_position text,
    role             text,
    created_at       timestamp,
    updated_at       timestamp,
    position_id      bigint not null
        constraint fk_positions_users
            references public.positions
);

alter table public.users
    owner to postgres;

create table if not exists public.shifts
(
    id          bigserial
        primary key,
    name        varchar(5),
    start_time  timestamp with time zone,
    end_time    timestamp with time zone,
    position_id bigserial
        constraint fk_positions_shifts
            references public.positions,
    user_id     bigserial
        constraint fk_users_shifts
            references public.users
);

alter table public.shifts
    owner to postgres;

