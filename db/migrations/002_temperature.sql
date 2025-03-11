-- create sensors type enum
CREATE TYPE sensor_type AS ENUM ('DHT11', 'DHT22');

-- create sensors table
CREATE TABLE sensors (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL REFERENCES users(id) NOT NULL,
  name TEXT NOT NULL,
  type sensor_type NOT NULL,
  mac_address VARCHAR(17) NOT NULL UNIQUE
);

-- create sensor data table
CREATE TABLE sensor_data (
  id BIGSERIAL PRIMARY KEY,
  sensor_id BIGSERIAL REFERENCES sensors(id) NOT NULL,
  temperature REAL NOT NULL,
  humidity REAL NOT NULL,
  reading_time TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

---- create above / drop below ----

DROP TABLE sensor_data;

DROP TABLE sensors;

DROP TYPE sensor_type;

