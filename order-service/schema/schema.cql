CREATE KEYSPACE IF NOT EXISTS ecommerce 
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE ecommerce;

CREATE TABLE IF NOT EXISTS orders (
    id uuid,
    product_id int,
    quantity int,
    total_price double,
    status text,
    created_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
    id int,
    name text,
    description text,
    price double,
    stock int,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders_by_product (
    product_id int,
    id uuid,
    quantity int,
    total_price double,
    status text,
    created_at timestamp,
    PRIMARY KEY (product_id, id)
);