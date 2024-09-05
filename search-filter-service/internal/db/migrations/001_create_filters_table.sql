CREATE TABLE IF NOT EXISTS filters (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    sender VARCHAR(255),
    subject_contains VARCHAR(255),
    move_to_folder VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
