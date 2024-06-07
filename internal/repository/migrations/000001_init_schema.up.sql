create database schedules_data
    with owner postgres;

create table public.positions
(
    id            bigserial
        primary key,
    position_name text
);

alter table public.positions
    owner to postgres;

create table public.users
(
    id            bigserial
        primary key,
    name_surname  text,
    email         text,
    password      text,
    position_name text,
    created_at    timestamp,
    updated_at    timestamp,
    user_id       bigint
        constraint fk_positions_users
            references public.positions
);

alter table public.users
    owner to postgres;

create table public.shifts
(
    id          bigserial
        primary key,
    name        varchar(15),
    start_time  time,
    end_time    time,
    position_id bigint
        constraint fk_positions_shifts
            references public.positions,
    user_id     bigint
        constraint fk_users_shifts
            references public.users
);

alter table public.shifts
    owner to postgres;

create table public.admins
(
    id        bigserial
        primary key,
    user_name varchar(255) not null,
    password  varchar(255) not null
);

alter table public.admins
    owner to postgres;

create table public.annual_leaves
(
    id            bigserial
        primary key,
    email         text,
    position_name text,
    start_date    date,
    end_date      date
);

alter table public.annual_leaves
    owner to postgres;

