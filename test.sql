CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    login VARCHAR,
    password VARCHAR
);

CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    order INTEGER,
    password VARCHAR
);
