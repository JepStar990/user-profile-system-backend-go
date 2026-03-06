CREATE TABLE IF NOT EXISTS user_history (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    content_id VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    last_position_seconds INT DEFAULT 0,
    duration_seconds INT DEFAULT 0,
    completed BOOLEAN DEFAULT FALSE,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY unique_history (user_id, content_id, content_type),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
