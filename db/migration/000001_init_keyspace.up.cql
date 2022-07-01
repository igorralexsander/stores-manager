CREATE TABLE stores (
    id uuid,
    name text,
    url text,
    max_timeout int,
    PRIMARY KEY (id, name)
);