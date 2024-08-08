-- create sensors type enum
CREATE TYPE sensor_type AS ENUM ('dht11', 'dht22');

-- create sensors table
CREATE TABLE sensors (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL REFERENCES users(id),
  name TEXT NOT NULL,
  type sensor_type NOT NULL,
  mac_address VARCHAR(17) NOT NULL UNIQUE
);

-- create sensor data table
CREATE TABLE sensor_data (
  id BIGSERIAL PRIMARY KEY,
  sensor_id BIGSERIAL REFERENCES sensors(id),
  temperature NUMERIC(5, 2),
  humidity NUMERIC(5, 2),
  reading_time TIMESTAMPTZ DEFAULT NOW()
);

---- create above / drop below ----

DROP TYPE sensor_type;

DROP TABLE sensors;

DROP TABLE sensor_data;

