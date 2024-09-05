CREATE TABLE IF NOT EXISTS emails (
    id SERIAL PRIMARY KEY,
    sender VARCHAR(100) NOT NULL,
    recipients TEXT[] NOT NULL,
    cc TEXT[],
    bcc TEXT[],
    subject VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    attachments TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_trashed BOOLEAN DEFAULT FALSE
);
