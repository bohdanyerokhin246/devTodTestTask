CREATE TABLE IF NOT EXISTS targets (
                                        id SERIAL PRIMARY KEY,
                                        mission_id INTEGER NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
                                        name VARCHAR(255) NOT NULL,
                                        country VARCHAR(100) NOT NULL,
                                        notes TEXT,
                                        is_complete BOOLEAN NOT NULL DEFAULT FALSE,
                                        created_at TIMESTAMP,
                                        updated_at TIMESTAMP,
                                        deleted_at TIMESTAMP
);