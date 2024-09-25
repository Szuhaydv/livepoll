SELECT votes FROM options
WHERE id = $1
FOR UPDATE;
