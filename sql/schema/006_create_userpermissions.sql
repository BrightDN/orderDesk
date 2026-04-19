-- +goose Up
CREATE TABLE user_permissions (
    user_id integer,
    permission_id integer,

    primary key (user_id, permission_id),

    FOREIGN KEY (user_id)
        REFERENCES users.(id)
        ON DELETE CASCADE,

    FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_permissions;