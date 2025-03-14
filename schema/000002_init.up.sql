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
    ALTER TABLE tracks ADD COLUMN done BOOLEAN DEFAULT FALSE;
    );

CREATE TABLE if not exists users
(
    name VARCHAR(255) NOT NULL,
    id serial PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    author_id INT REFERENCES authors(id) ON DELETE SET NULL
    );

CREATE TABLE IF NOT EXISTS users_albums (
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    album_id INT NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, album_id)
    );
