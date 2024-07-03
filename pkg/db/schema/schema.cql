CREATE KEYSPACE todo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE todo.users (
    user_id timeuuid PRIMARY KEY,
    first_name text,
    last_name text
);

CREATE TABLE todo.items (
    user_id timeuuid,
    id timeuuid,
    created_at timestamp,
    description text,
    status boolean,
    title text,
    updated_at timestamp,
    PRIMARY KEY ((user_id), id)
);