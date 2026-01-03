\connect surf-share;

CREATE SCHEMA IF NOT EXISTS app;

CREATE TABLE IF NOT EXISTS app.breaks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    coordinates POINT NOT NULL,
    country VARCHAR(3) NOT NULL,
    region VARCHAR(100),
    city VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COPY app.breaks(name,slug,description,coordinates,country,region,city)
  FROM '/docker-entrypoint-initdb.d/breaks_seeds.csv' CSV HEADER;