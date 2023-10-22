-- Adding a status enum type
CREATE TYPE status_type AS ENUM('pending', 'in-progress', 'completed');

-- Adding a table for order
CREATE TABLE IF NOT EXISTS "order" (
    id SERIAL PRIMARY KEY,
    notes VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    menu_id INT,
    user_id INT,
    status status_type DEFAULT 'pending' NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT true,
    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE SET NULL,
    CONSTRAINT fk_menu_id
    FOREIGN KEY(menu_id)
    REFERENCES menu(id)
    ON DELETE SET NULL
);
