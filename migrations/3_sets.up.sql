CREATE TABLE IF NOT EXISTS sets
(
    id SERIAL PRIMARY KEY,
    sport_id INTEGER REFERENCES sports(id),
    brand_id INTEGER REFERENCES brands(id),
    year INTEGER NOT NULL,
    set VARCHAR NOT NULL
)