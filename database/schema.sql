CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL ,
    password VARCHAR(255) NOT NULL,
    UNIQUE(email),
    UNIQUE(username)
);

CREATE TABLE books(
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    sellerName VARCHAR(255) NOT NULL,
    condition BOOLEAN NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_seller
        FOREIGN KEY(sellerName)
            REFERENCES users(username)
)