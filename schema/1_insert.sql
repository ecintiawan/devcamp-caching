INSERT INTO items
SELECT id, 'alfa'
FROM generate_series(1, 10000) AS id;