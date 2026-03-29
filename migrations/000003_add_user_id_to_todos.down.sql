-- Drop foreign key constraint first
ALTER TABLE todos
DROP CONSTRAINT IF EXISTS fk_todos_user_id;

-- Drop the column
ALTER TABLE todos
DROP COLUMN user_id;