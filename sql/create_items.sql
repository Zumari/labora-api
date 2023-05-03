CREATE DATABASE labora_proyect_1;

CREATE TABLE items (
    id int PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    order_date DATE NOT NULL,
	product VARCHAR NOT NULL,
	quantity INTEGER NOT NULL,
	price NUMERIC NOT NULL
);