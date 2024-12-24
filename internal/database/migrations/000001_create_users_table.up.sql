CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    login VARCHAR(30) UNIQUE,
    password VARCHAR(30),
    balance REAL DEFAULT 0
);