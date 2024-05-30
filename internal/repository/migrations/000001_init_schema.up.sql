create table if not exists public.admins
(
    id        bigint       not null
        primary key,
    user_name varchar(255) not null,
    password  varchar(255) not null
);

alter table public.admins
    owner to postgres;

create table if not exists public.positions
(
    id            bigint not null
        primary key,
    position_name text
);

alter table public.positions
    owner to postgres;

create table if not exists public.users
(
    id            bigint not null
        primary key,
    name_surname  text,
    email         text,
    password      text,
    position_name text,
    created_at    timestamp,
    updated_at    timestamp,
    position_id   bigint
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
    position_id bigint
        constraint fk_positions_shifts
            references public.positions,
    user_id     bigint
        constraint fk_users_shifts
            references public.users
);

alter table public.shifts
    owner to postgres;