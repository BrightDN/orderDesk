-- +goose Up

INSERT INTO roles (name) VALUES
    ('site_admin'),
    ('superadmin'),
    ('admin'),
    ('employee'),
    ('custom');

-- +goose Down

DELETE FROM roles;