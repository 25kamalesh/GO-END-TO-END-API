-- Add user_id column (nullable initially)
ALTER TABLE todos
ADD COLUMN user_id INT;

-- Update existing rows to use the first user (or NULL if no users exist)
-- You may want to change this to a specific user_id
UPDATE todos
SET user_id = (SELECT id FROM users ORDER BY id LIMIT 1)
WHERE user_id IS NULL;

-- Now make it NOT NULL
ALTER TABLE todos
ALTER COLUMN user_id SET NOT NULL;

-- Add foreign key constraint
ALTER TABLE todos
ADD CONSTRAINT fk_todos_user_id
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;