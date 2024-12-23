CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    login VARCHAR(30) UNIQUE,
    password VARCHAR(30),
    balance INTEGER
);

CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    user_id FOREIGN KEY,
    order INTEGER UNIQUE,
    upload_at DATE
);

CREATE TABLE Withdrawals(
    id SERIAL PRIMARY KEY,
    user_id FOREIGN KEY,
    order_id FOREIGN KEY,
    processed_at DATE
)