-- Adding a table for menu
CREATE TABLE IF NOT EXISTS menu (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    price INT NOT NULL,
    cafe_id INT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT true,
    CONSTRAINT fk_cafe_id
    FOREIGN KEY(cafe_id)
    REFERENCES cafe(id)
    ON DELETE SET NULL
);
