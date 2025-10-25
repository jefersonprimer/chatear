-- Triggers
CREATE TRIGGER enforce_bucket_name_length_trigger
BEFORE INSERT OR UPDATE ON storage.buckets
FOR EACH ROW EXECUTE FUNCTION storage.enforce_bucket_name_length();

CREATE TRIGGER objects_delete_delete_prefix
AFTER DELETE ON storage.objects
FOR EACH ROW EXECUTE FUNCTION storage.delete_prefix_hierarchy_trigger();

CREATE TRIGGER objects_insert_create_prefix
BEFORE INSERT ON storage.objects
FOR EACH ROW EXECUTE FUNCTION storage.objects_insert_prefix_trigger();

CREATE TRIGGER objects_update_create_prefix
BEFORE UPDATE ON storage.objects
FOR EACH ROW EXECUTE FUNCTION storage.objects_update_prefix_trigger();

CREATE TRIGGER update_objects_updated_at
BEFORE UPDATE ON storage.objects
FOR EACH ROW EXECUTE FUNCTION storage.update_updated_at_column();

CREATE TRIGGER prefixes_create_hierarchy
BEFORE INSERT ON storage.prefixes
FOR EACH ROW EXECUTE FUNCTION storage.prefixes_insert_trigger();

CREATE TRIGGER prefixes_delete_hierarchy
AFTER DELETE ON storage.prefixes
FOR EACH ROW EXECUTE FUNCTION storage.delete_prefix_hierarchy_trigger();

CREATE TRIGGER tr_check_filters
BEFORE INSERT OR UPDATE ON realtime.subscription
FOR EACH ROW EXECUTE FUNCTION realtime.subscription_check_filters();
