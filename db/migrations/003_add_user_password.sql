-- Write your migrate up statements here

ALTER TABLE users
ADD COLUMN password BYTEA NOT NULL;

---- create above / drop below ----

ALTER TABLE users
DROP COLUMN password;

