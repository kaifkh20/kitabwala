CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL ,
    password VARCHAR(255) NOT NULL,
    UNIQUE(email),
    UNIQUE(username)
);


CREATE TABLE IF NOT EXISTS books(
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    sellerName VARCHAR(255) NOT NULL,
    condition BOOLEAN NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_seller
        FOREIGN KEY(sellerName)
            REFERENCES users(email)
);

CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL,
    userId BIGSERIAL,
    bookId BIGSERIAL,
    PRIMARY KEY(id),
    FOREIGN KEY(userId) REFERENCES users(id),
    FOREIGN KEY(bookId) REFERENCES books(id)
);