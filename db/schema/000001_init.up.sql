create table users(
    id serial primary key not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);