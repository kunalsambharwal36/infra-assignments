CREATE TABLE IF NOT EXISTS configs (
    id VARCHAR(255) PRIMARY KEY,
    host VARCHAR(255),
    port INT,
    app_name VARCHAR(255),
    log_level VARCHAR(50)
);