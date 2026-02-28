create table public.admins
(
    id        bigint       not null
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


alter table public.annual_leaves
    owner to postgres;

create table public.daily_schedules
(
    id              bigserial
        primary key,
    start_date      timestamp with time zone,
    positions_names text[],
    employees       text[],
    shifts          text[]
);

create table public.daily_assignments
(
    id            bigserial
        primary key,
    sender_email  text,
    position_name text,
    description   text,
    sent_count    bigint,
    created_at    timestamp
);

alter table public.daily_assignments
    owner to postgres;

create index idx_daily_assignments_position_name
    on public.daily_assignments (position_name);

create index idx_daily_assignments_sender_email
    on public.daily_assignments (sender_email);

create table public.daily_schedules
(
    id              bigserial
        primary key,
    start_date      timestamp with time zone,
    positions_names text[],
    employees       text[],
    shifts          text[]
);

alter table public.daily_schedules
    owner to postgres;


create table public.daily_schedule_positions
(
    daily_schedule_id bigint not null
        constraint fk_daily_schedule_positions_daily_schedule
            references public.daily_schedules,
    positions_id      bigint not null
        constraint fk_daily_schedule_positions_positions
            references public.positions,
    primary key (daily_schedule_id, positions_id)
);

alter table public.daily_schedule_positions
    owner to postgres;



alter table public.daily_schedules
    owner to postgres;

create table public.positions
(
    id            bigserial
        primary key,
    position_name text,
    position_id   bigint
        constraint fk_daily_schedules_positions
            references public.daily_schedules
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

create index idx_users_position_name
    on public.users (position_name);

create table public.shifts
(
    id          bigint not null
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

create index idx_shifts_position_id
    on public.shifts (position_id);

create index idx_shifts_user_id
    on public.shifts (user_id);

