create table roles
(
    role_id serial primary key,
    name varchar(31) not null unique,
    description varchar(127) not null,
    significance_order int not null
);
create table users
(
    user_id serial primary key,
    login varchar(63) not null unique,
    password varchar(63) NOT NULL,
    nickname varchar(63) not null,
    email varchar(127) not null unique,
    sign_up_date timestamp not null,
    role_id int not null references roles(role_id)
);

CREATE TABLE dumps (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL
);