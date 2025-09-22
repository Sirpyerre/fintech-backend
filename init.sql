-- init.sql
-- Create users table
CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY,
    name       VARCHAR(50)              NOT NULL,
    lastname   VARCHAR(50)              NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_id CHECK (id > 0)
);

-- Create transactions table with foreign key to users
CREATE TABLE IF NOT EXISTS transactions
(
    id       SERIAL PRIMARY KEY,
    user_id  INTEGER                  NOT NULL,
    amount   NUMERIC(10, 2)           NOT NULL,
    datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT unique_id UNIQUE (id),
    CONSTRAINT positive_user_id CHECK ( user_id > 0 ),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE RESTRICT
);

-- Indexes for efficient querying
CREATE INDEX idx_transactions_user_id ON transactions (user_id);
CREATE INDEX idx_transactions_datetime ON transactions (datetime);

-- Insert sample data into users table
INSERT INTO users (id, name, lastname, created_at)
VALUES (1, 'Peter', 'Doe', '2024-06-01T10:00:00Z'),
       (2, 'Jane', 'Smith', '2024-06-01T10:00:00Z'),
       (3, 'John', 'Brown', '2024-06-01T10:00:00Z'),
       (4, 'Alice', 'Johnson', '2024-06-01T10:00:00Z'),
       (5, 'Bob', 'Davis', '2024-06-01T10:00:00Z'),
       (6, 'Charlie', 'Miller', '2024-06-01T10:00:00Z'),
       (7, 'Diana', 'Wilson', '2024-06-01T10:00:00Z'),
       (8, 'Ethan', 'Moore', '2024-06-01T10:00:00Z'),
       (9, 'Fiona', 'Taylor', '2024-06-01T10:00:00Z'),
       (10, 'George', 'Anderson', '2024-06-01T10:00:00Z');