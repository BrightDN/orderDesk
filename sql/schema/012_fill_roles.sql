-- +goose Up

INSERT INTO roles (name) VALUES
    ('superadmin'),
    ('admin'),
    ('employee'),
    ('custom');

-- +goose Down

DELETE FROM roles;