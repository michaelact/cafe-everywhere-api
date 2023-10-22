-- Adding an table for cafe
CREATE TABLE IF NOT EXISTS cafe (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Adding an index for the email column
CREATE INDEX idx_cafe_email ON cafe (email);