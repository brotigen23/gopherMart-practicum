CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    login VARCHAR(30) UNIQUE,
    password VARCHAR(30),
    balance REAL
);

CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users (id),
    order INTEGER UNIQUE,
    uploaded_at DATE
);

CREATE TABLE Withdrawals(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users (id),
    sum REAL,
    processed_at DATE
)