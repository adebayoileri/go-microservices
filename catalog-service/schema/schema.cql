CREATE KEYSPACE IF NOT EXISTS ecommerce 
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE ecommerce;

DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS products (
    id int,
    name text,
    description text,
    price double,
    stock int,
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS products_name_idx ON products(name); 
