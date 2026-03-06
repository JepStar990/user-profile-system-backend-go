CREATE TABLE IF NOT EXISTS user_favorites (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    content_id VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    title VARCHAR(255),
    preview TEXT,
    duration_seconds INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    UNIQUE KEY unique_favorite (user_id, content_id, content_type),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
