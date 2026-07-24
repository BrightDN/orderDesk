-- +goose Up

CREATE TABLE employee_permissions (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER,
    permission_id INTEGER,

    UNIQUE (employee_id, permission_id),

    FOREIGN KEY (employee_id)
        REFERENCES employees(id)
        ON DELETE CASCADE,

    FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE CASCADE  

);

-- +goose Down

DROP TABLE IF EXISTS employee_permissions;