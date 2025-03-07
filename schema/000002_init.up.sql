CREATE TABLE if not exists authors
(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
    );

CREATE TABLE if not exists albums
(
    id serial PRIMARY KEY,
    author_id INT NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE if not exists tracks
(
    id serial PRIMARY KEY,
    album_id INT NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    duration INT CHECK (duration > 0)
    );

