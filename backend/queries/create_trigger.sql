CREATE TRIGGER new_vote_trigger
AFTER INSERT OR UPDATE ON options -- maybe add DELETE later
FOR EACH ROW EXECUTE FUNCTION notify_new_vote();
