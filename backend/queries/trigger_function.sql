CREATE OR REPLACE FUNCTION notify_new_vote()
RETURNS TRIGGER AS $$
DECLARE
    poll_id UUID;
    votes_count INT;
BEGIN
    SELECT poll_id INTO poll_id FROM options WHERE id = NEW.id;
    SELECT votes INTO votes_count FROM options WHERE id = NEW.id;

    PERFORM pg_notify('poll_' || poll_id, json_build_object('option_id', NEW.id, 'votes', votes_count)::text);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
