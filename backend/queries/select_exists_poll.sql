SELECT EXISTS (
    SELECT 1
    FROM polls
    WHERE id = $1
);
