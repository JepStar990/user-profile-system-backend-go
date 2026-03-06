CREATE TABLE IF NOT EXISTS user_listening_events (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    content_id VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    session_id CHAR(36) NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    position_seconds INT DEFAULT 0,
    duration_seconds INT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_user (user_id),
    INDEX idx_session (session_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
