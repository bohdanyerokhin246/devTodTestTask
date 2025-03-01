CREATE TABLE IF NOT EXISTS missions (
                                    id SERIAL PRIMARY KEY,
                                    cat_id INTEGER NOT NULL REFERENCES cats(id) ON DELETE CASCADE,
                                    is_complete BOOLEAN NOT NULL DEFAULT FALSE,
                                    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                    deleted_at TIMESTAMP NULL
);