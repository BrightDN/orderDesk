-- +goose Up

INSERT INTO roles (name) {
    VALUES
    ('superadmin'),
    ('admin'),
    ('employee'),
    ('custom')
};

-- +goose Down

TRUNCATE TABLE roles;