CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    balance INTEGER,
    surname VARCHAR(64),
    username VARCHAR(64),
    password VARCHAR (64)
)
