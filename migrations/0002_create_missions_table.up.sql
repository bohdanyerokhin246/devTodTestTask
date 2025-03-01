CREATE TABLE IF NOT EXISTS missions (
                                    id SERIAL PRIMARY KEY,
                                    cat_id INTEGER NOT NULL REFERENCES cats(id) ON DELETE CASCADE,
                                    is_complete BOOLEAN NOT NULL DEFAULT FALSE,
                                    created_at TIMESTAMP,
                                    updated_at TIMESTAMP,
                                    deleted_at TIMESTAMP
);