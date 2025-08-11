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


CREATE OR REPLACE FUNCTION post_user(
    p_login VARCHAR(63),
    p_email VARCHAR(127),
    p_password VARCHAR(63),
    p_nickname VARCHAR(63),
    p_registered TIMESTAMP,
    p_role_id INT
) RETURNS INT AS $$
DECLARE
    new_user_id INT;
BEGIN
    INSERT INTO users
    VALUES (p_login, p_email, p_password, p_nickname, p_registered, p_role_id)
    RETURNING id INTO new_user_id;

    RETURN new_user_id;
END;
$$ LANGUAGE plpgsql;