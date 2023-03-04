CREATE DATABASE word_proximity;

\ c word_proximity;

CREATE TABLE IF NOT EXISTS words (
    id SERIAL PRIMARY KEY,
    word TEXT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_uniq_words_word ON words (word);

CREATE UNIQUE INDEX idx_uniq_words_date ON words (date);

INSERT INTO
    words (word, date)
VALUES
    ('APPLE', NOW());

CREATE TABLE IF NOT EXISTS allowed_words (
    id SERIAL PRIMARY KEY,
    word TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_uniq_allowed_words_word ON allowed_words (word);

ALTER TABLE
    ONLY words
ADD
    CONSTRAINT words_word_fkey FOREIGN KEY (word) REFERENCES allowed_words(word);