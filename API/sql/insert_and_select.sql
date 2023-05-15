
DELETE items;
DROP TABLE items;

INSERT INTO public.items(
	customer_name, order_date, product, quantity, price)
	VALUES ('Samantha', '23/05/03 16:06:41.50', 'Memoria USB 1TB', 25, 150.00),
			('Pedro', '23/05/03 16:06:43.50', 'Vino La Rosa', 2, 120.00),
			('Selim', '23/05/03 16:07:41.50', 'Caja plástica', 205, 100.00),
			('Gloria', '23/05/03 16:10:41.50', 'Cepillo dental', 1, 24.50),
			('Alejandra', '23/05/03 16:16:00.50', 'Mouse', 2, 250.49)
;
	
SELECT * FROM items
WHERE quantity > 2 AND price > 50;

