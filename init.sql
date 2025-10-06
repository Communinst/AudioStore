
create table roles
(
    id smallint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    significance_order int not null,
    alias varchar(31) not null unique
);


create table users
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login varchar(63) not null unique,
    email varchar(127) not null unique,
    password varchar(63) NOT NULL,
    nickname varchar(63) not null,
    registered timestamp not null,
    role_id smallint not null references roles(id)  
);


CREATE TABLE dumps (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
    filename VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL
);


-- CREATE OR REPLACE FUNCTION insert_user(
--     p_login VARCHAR(63),
--     p_email VARCHAR(127),
--     p_password VARCHAR(63),
--     p_nickname VARCHAR(63),
--     p_registered TIMESTAMP,
--     p_role_id SMALLINT  
-- ) RETURNS BIGINT AS $$  
-- DECLARE
--     new_id BIGINT;
-- BEGIN
--     INSERT INTO users (login, email, password, nickname, registered, role_id)
--     VALUES (p_login, p_email, p_password, p_nickname, p_registered, p_role_id)
--     RETURNING id INTO new_id;
    
--     RETURN new_id;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION get_user_by_id(p_id BIGINT)  
-- RETURNS TABLE (
--     id BIGINT,
--     login VARCHAR(63),
--     email VARCHAR(127),
--     password VARCHAR(63),
--     nickname VARCHAR(63),
--     registered TIMESTAMP,
--     role_id SMALLINT  
-- ) AS $$
-- BEGIN
--     RETURN QUERY
--     SELECT *
--     FROM users u
--     WHERE u.id = p_id;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION get_users_by_ids(p_ids BIGINT[])  
-- RETURNS TABLE (
--     id BIGINT,
--     login VARCHAR(63),
--     email VARCHAR(127),
--     password VARCHAR(63),
--     nickname VARCHAR(63),
--     registered TIMESTAMP,
--     role_id SMALLINT  
-- ) AS $$
-- BEGIN
--     RETURN QUERY
--     SELECT u.id, u.login, u.email, u.password, u.nickname, u.registered, u.role_id
--     FROM users u
--     WHERE u.id = ANY(p_ids);
-- END;
-- $$ LANGUAGE plpgsql;


-- CREATE OR REPLACE FUNCTION get_all_users() 
-- RETURNS TABLE (
--     id BIGINT,
--     login VARCHAR(63),
--     email VARCHAR(127),
--     password VARCHAR(63),
--     nickname VARCHAR(63),
--     registered TIMESTAMP,
--     role_id SMALLINT
-- ) AS $$
-- BEGIN
--     RETURN QUERY
--     SELECT u.id, u.login, u.email, u.password, u.nickname, u.registered, u.role_id
--     FROM users u
--     ORDER BY u.id;  
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION delete_user_by_id(p_id BIGINT) 
-- RETURNS BIGINT AS $$
-- DECLARE
--     deleted_count BIGINT;
-- BEGIN
--     DELETE FROM users WHERE id = p_id;
--     GET DIAGNOSTICS deleted_count = ROW_COUNT;
--     RETURN deleted_count;
-- END;
-- $$ LANGUAGE plpgsql;


-- CREATE OR REPLACE FUNCTION delete_users_by_ids(p_ids BIGINT[]) 
-- RETURNS VOID AS $$
-- DECLARE
--     deleted_count BIGINT;
-- BEGIN
--     DELETE FROM users WHERE id = ANY(p_ids);
--     GET DIAGNOSTICS deleted_count = ROW_COUNT;
--     RETURN deleted_count;
-- END;
-- $$ LANGUAGE plpgsql;

