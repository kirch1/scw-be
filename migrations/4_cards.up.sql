CREATE TABLE IF NOT EXISTS cards
(
    id SERIAL PRIMARY KEY,
    set_id INTEGER REFERENCES sets(id),
    athlete VARCHAR,
    team VARCHAR,
    position VARCHAR,
    num VARCHAR,
    sequence VARCHAR
)