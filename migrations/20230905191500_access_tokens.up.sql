CREATE TABLE IF NOT EXISTS access_tokens (
                                             id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                             user_id BIGINT UNSIGNED NOT NULL,
                                             ip_address VARCHAR(255),
    user_agent VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    expire_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id)
    REFERENCES users(id)
                                                   ON DELETE CASCADE
                                                   ON UPDATE CASCADE
    );