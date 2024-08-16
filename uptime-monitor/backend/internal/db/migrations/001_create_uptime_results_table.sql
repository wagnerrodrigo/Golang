CREATE TABLE uptime_results (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    status BOOLEAN NOT NULL,
    checked_at TIMESTAMP NOT NULL
);