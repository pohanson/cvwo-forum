DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "appuser";
DROP TABLE IF EXISTS "thread_type";
DROP TABLE IF EXISTS "thread";
DROP TABLE IF EXISTS "thread_jt";
DROP TABLE IF EXISTS "category";
DROP TABLE IF EXISTS "thread_category";
DROP TRIGGER IF EXISTS updating_edited_on_thread ON "thread";
DROP FUNCTION IF EXISTS update_edited_time;
