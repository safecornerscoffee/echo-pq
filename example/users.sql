CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);

INSERT INTO users (age, email, first_name, last_name)
VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun');
INSERT INTO users (age, email, first_name, last_name)
VALUES (52, 'bob@smith.io', 'Bob', 'Smith');
INSERT INTO users (age, email, first_name, last_name)
VALUES (15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');
