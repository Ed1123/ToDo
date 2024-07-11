-- Create todos table
DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    archived BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create tasks table
DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    date_created TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    date_completed TIMESTAMP,
    archived BOOLEAN NOT NULL DEFAULT FALSE,
    todo_id INTEGER REFERENCES todos(id)
);

-- Create task history table
DROP TABLE IF EXISTS task_history;
CREATE TABLE task_history (
    id SERIAL PRIMARY KEY,
    task_id INTEGER REFERENCES tasks(id),
    date TIMESTAMP NOT NULL,
    todo_id INTEGER REFERENCES todos(id)
);
