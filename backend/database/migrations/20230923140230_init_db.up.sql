create table if not exists public.hotels
(
    id              bigserial
        primary key,
    name            text,
    stars           bigint,
    rating          numeric,
    count_reviews   bigint,
    amenities       text,
    address         text,
    price_per_night bigint,
    url_image       text,
    freebies        text
);

alter table public.hotels
    owner to postgres;

create table if not exists public.rooms
(
    id                 bigserial
        primary key,
    hotel_id           bigint
        constraint fk_hotels_rooms
            references public.hotels
            on delete cascade,
    type               text,
    is_view_on_city    boolean,
    amount_double_beds bigint,
    amount_single_beds bigint,
    cost_per_night     bigint
);

alter table public.rooms
    owner to postgres;

create table if not exists public.users
(
    id            bigserial
        primary key,
    first_name    text,
    last_name     text,
    email         text,
    password      text,
    phone_number  text,
    address       text,
    date_of_birth text
);

alter table public.users
    owner to postgres;

create table if not exists public.history_bookings
(
    id      bigserial
        primary key,
    user_id bigint
        constraint fk_users_history_booking
            references public.users
            on delete cascade
);

alter table public.history_bookings
    owner to postgres;

create table if not exists public.bookings
(
    id             bigserial
        primary key,
    booking_id     bigint
        constraint fk_booking_id
            references public.history_bookings
            on delete cascade,
    room_id        bigint,
    check_in_time  timestamp,
    check_out_time timestamp,
    createdat      timestamp with time zone default CURRENT_TIMESTAMP,
    updatedat      timestamp with time zone default CURRENT_TIMESTAMP
);

alter table public.bookings
    owner to postgres;

