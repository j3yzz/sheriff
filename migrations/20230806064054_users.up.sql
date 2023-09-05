CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(12) NOT NULL UNIQUE,
    phone_verified_at TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);