DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_trigger
        WHERE tgname = 'new_vote_trigger'
    ) THEN
        CREATE TRIGGER new_vote_trigger
        AFTER INSERT OR UPDATE ON options -- maybe add DELETE later
        FOR EACH ROW EXECUTE FUNCTION notify_new_vote();
    END IF;
END;
$$;
