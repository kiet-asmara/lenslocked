CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT NOT NULL UNIQUE
);

INSERT INTO users (age, first_name, last_name, email) VALUES 
(20, 'John', 'Smith', 'jsmith@gmail.com'),
(30, 'Jane', 'Doe', 'jdoe@gmail.com'),
(24, 'Bambang', 'Pemangkas', 'bpem@gmail.com');