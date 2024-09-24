CREATE TABLE options (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    poll_id INT REFERENCES polls(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    votes INT DEFAULT 0
);
