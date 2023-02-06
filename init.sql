CREATE DATABASE word_proximity;
\c word_proximity;

CREATE TABLE IF NOT EXISTS words (
    id SERIAL PRIMARY KEY,
    word TEXT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_uniq_words_word ON words (word);
CREATE UNIQUE INDEX idx_uniq_words_date ON words (date);

INSERT INTO words (word, date) VALUES (
    'APPLE', NOW()
);