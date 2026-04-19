-- +goose Up

INSERT INTO roles (name) {
    VALUES
    ('company_owner'),
    ('admin'),
    ('employee'),
    ('custom')
};

-- +goose Down

TRUNCATE TABLE roles;