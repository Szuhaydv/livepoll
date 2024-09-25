INSERT INTO polls (id, duration, title)
VALUES ($1, COALESCE($2, DEFAULT), $3);
