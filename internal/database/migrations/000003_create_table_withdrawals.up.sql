CREATE TABLE Withdrawals(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    sum REAL,
    processed_at TIMESTAMP WITH TIME ZONE
);