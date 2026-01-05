\connect surf-share;

CREATE SCHEMA IF NOT EXISTS app;

BEGIN;

CREATE TABLE IF NOT EXISTS app.breaks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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

CREATE TABLE IF NOT EXISTS app.breaks_media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    break_slug VARCHAR(100) NOT NULL REFERENCES app.breaks(slug) ON DELETE CASCADE,
    video_url TEXT,
    image_urls TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COPY app.breaks_media(break_slug,video_url,image_urls)
    FROM '/docker-entrypoint-initdb.d/breaks_media_seeds.csv' CSV HEADER;

COMMIT;