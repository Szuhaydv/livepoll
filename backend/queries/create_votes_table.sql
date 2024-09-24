CREATE TABLE votes (
    voter_id INET PRIMARY KEY,
    option_id UUID REFERENCES options(id) ON DELETE CASCADE,
)
