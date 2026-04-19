-- SCHEMA: NOT USED
-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    ID INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1),
    User_UUID uuid DEFAULT gen_random_uuid(),
    User_Email TEXT NOT NULL UNIQUE,
    First_Name TEXT NOT NULL,
    Last_Name TEXT
);

-- Create a default user for now
INSERT INTO users(User_Email, First_Name, Last_Name)
VALUES
    ('a.lab@acciaiolab.com', 'Acciaio', 'Lab'),
    ('demo@acciaiolab.com', 'Demo', 'User'),
    ('test@acciaiolab.com', 'Test', 'User');