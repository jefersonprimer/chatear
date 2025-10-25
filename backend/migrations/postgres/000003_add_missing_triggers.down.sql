-- Triggers
DROP TRIGGER enforce_bucket_name_length_trigger ON storage.buckets;
DROP TRIGGER objects_delete_delete_prefix ON storage.objects;
DROP TRIGGER objects_insert_create_prefix ON storage.objects;
DROP TRIGGER objects_update_create_prefix ON storage.objects;
DROP TRIGGER update_objects_updated_at ON storage.objects;
DROP TRIGGER prefixes_create_hierarchy ON storage.prefixes;
DROP TRIGGER prefixes_delete_hierarchy ON storage.prefixes;
DROP TRIGGER tr_check_filters ON realtime.subscription;
