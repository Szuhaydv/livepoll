CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    voter_id INET NOT NULL,
    option_id UUID REFERENCES options(id) ON DELETE CASCADE,
    poll_id UUID REFERENCES polls(id) ON DELETE CASCADE,
    CONSTRAINT unique_vote_per_poll UNIQUE (poll_id, voter_id)
);
