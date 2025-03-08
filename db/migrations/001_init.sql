-- create users table
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL
);

---- create above / drop below ----

DROP TABLE users;
