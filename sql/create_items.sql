CREATE DATABASE labora-proyect-1;

CREATE TABLE items(
	id SERIAL PRIMARY KEY,
	customer_name VARCHAR(255) NOT NULL,
	order_date TIMESTAMP NOT NULL,
	product VARCHAR(255) NOT NULL,
	quantity INTEGER NOT NULL,
	price NUMERIC NOT NULL
);