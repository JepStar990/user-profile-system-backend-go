CREATE TABLE IF NOT EXISTS user_downloads (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    content_id VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    download_quality VARCHAR(50) NOT NULL,
    file_size_bytes BIGINT DEFAULT 0,
    storage_url VARCHAR(512) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'ready',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    UNIQUE KEY unique_download (user_id, content_id, content_type),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
