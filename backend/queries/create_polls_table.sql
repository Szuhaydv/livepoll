CREATE TABLE IF NOT EXISTS polls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    duration INTERVAL DEFAULT '120 second',
    title TEXT NOT NULL
);
