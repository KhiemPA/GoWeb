create database test;

\c test;

create table if not exists category(
    id serial PRIMARY key,
    description varchar(100) not null
);

create if not exists products (
    id bigserial PRIMARY KEY,
    name varchar(255) not null,
    price real not null,
    quantity integer default 0,
    amount real default 0.0,
    category bigint not null,
    CONSTRAINT products_category_fk foreign key(category)
    REFERENCES category(id)
);


create table if not exists users (
    id bigserial PRIMARY KEY,
    firtname varchar(15) NOT NULL,
    lastname varchar(20) NOT NULL,
    email varchar(40) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    status char(1) DEFAULT '0'
)