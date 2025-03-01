CREATE TABLE IF NOT EXISTS cats (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(100) NOT NULL,
                                    experience INT NOT NULL CHECK (experience >= 0),
                                    breed VARCHAR(100) NOT NULL,
                                    salary NUMERIC(10,2) NOT NULL CHECK (salary >= 0),
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);