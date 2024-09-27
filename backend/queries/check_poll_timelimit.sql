SELECT created_at, duration
FROM polls
WHERE id = $1;
