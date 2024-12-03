-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(100) NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(1000) NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down

DROP TABLE posts;
