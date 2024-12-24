CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    "order" VARCHAR UNIQUE,
    uploaded_at TIMESTAMP WITH TIME ZONE
);