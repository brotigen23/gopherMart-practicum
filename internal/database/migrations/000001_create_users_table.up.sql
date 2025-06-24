CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE,
    password TEXT,
    balance REAL DEFAULT 0
);