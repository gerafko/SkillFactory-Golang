-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO posts (id, author_id, title, content, created_at, updated_at) VALUES (0, 0, 'Статья', 'Содержание статьи', now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts, authors;
-- +goose StatementEnd