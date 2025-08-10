create table roles
(
    id smallserial primary key,
    order int not null
    alias varchar(31) not null unique,
);
create table users
(
    id serial primary key,
    login varchar(63) not null unique,
    email varchar(127) not null unique,
    password varchar(63) NOT NULL,
    nickname varchar(63) not null,
    registered timestamp not null,
    role_id int not null references roles(role_id)
);

CREATE TABLE dumps (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL
);
